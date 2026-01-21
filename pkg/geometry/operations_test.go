package geometry

import (
	"math"
	mrand "math/rand"
	"reflect"
	"testing"
)

func TestCompressCentroids_UnchangedWhenMaxNonPositiveOrLenLE(t *testing.T) {
	input := [][2]float64{{0, 0}, {1, 1}, {2, 2}}

	// maxPoints <= 0 returns input as-is
	if got := CompressCentroids(input, 0); !reflect.DeepEqual(got, input) {
		t.Fatalf("maxPoints=0: got %v, want %v", got, input)
	}

	if got := CompressCentroids(input, -5); !reflect.DeepEqual(got, input) {
		t.Fatalf("maxPoints=-5: got %v, want %v", got, input)
	}

	// len(centroids) <= maxPoints returns input as-is
	if got := CompressCentroids(input, len(input)); !reflect.DeepEqual(got, input) {
		t.Fatalf("maxPoints=len: got %v, want %v", got, input)
	}

	if got := CompressCentroids(input, len(input)+1); !reflect.DeepEqual(got, input) {
		t.Fatalf("maxPoints>len: got %v, want %v", got, input)
	}
}

func TestCompressCentroids_AllIdentical(t *testing.T) {
	input := [][2]float64{{1, 1}, {1, 1}, {1, 1}, {1, 1}}
	got := CompressCentroids(input, 2)
	if len(got) != 1 {
		t.Fatalf("identical points: got len=%d, want 1, got=%v", len(got), got)
	}
	if got[0] != input[0] {
		t.Fatalf("identical points: first mismatch: got %v, want %v", got[0], input[0])
	}
}

func TestCompressCentroids_DownsampleLine(t *testing.T) {
	// Points along a line x in [0,10], y=0; many points vs maxPoints to force downsampling
	const M = 101
	const maxPoints = 10
	centroids := make([][2]float64, 0, M)
	for i := 0; i < M; i++ {
		x := 10.0 * float64(i) / float64(M-1)
		centroids = append(centroids, [2]float64{x, 0})
	}

	got := CompressCentroids(centroids, maxPoints)

	if len(got) != maxPoints {
		t.Fatalf("downsample line: got len=%d, want %d; got=%v", len(got), maxPoints, got)
	}

	// Expect the first and last to be preserved after downsampling path
	if got[0] != centroids[0] {
		t.Fatalf("downsample line: first point mismatch: got %v, want %v", got[0], centroids[0])
	}
	if got[len(got)-1] != ([2]float64{10, 0}) {
		t.Fatalf("downsample line: last point mismatch: got %v, want %v", got[len(got)-1], [2]float64{10, 0})
	}

	// All returned points should be from the original set
	// Simple containment check by value (since we generated exact floats)
	contains := func(p [2]float64) bool {
		for _, q := range centroids {
			if q == p {
				return true
			}
		}
		return false
	}
	for _, p := range got {
		if !contains(p) {
			t.Fatalf("downsample line: result contains point not in input: %v", p)
		}
	}
}

func TestCompressCentroids_VerticalLineFallback(t *testing.T) {
	// width=0, height>0 ensures fallback cellSize path
	const M = 101
	const maxPoints = 7
	x := 5.0
	centroids := make([][2]float64, 0, M)
	for i := 0; i < M; i++ {
		y := 100.0 * float64(i) / float64(M-1)
		centroids = append(centroids, [2]float64{x, y})
	}

	got := CompressCentroids(centroids, maxPoints)
	if len(got) != maxPoints {
		t.Fatalf("vertical line: got len=%d, want %d; got=%v", len(got), maxPoints, got)
	}
	if got[0] != centroids[0] {
		t.Fatalf("vertical line: first point mismatch: got %v, want %v", got[0], centroids[0])
	}
	if got[len(got)-1] != centroids[len(centroids)-1] {
		t.Fatalf("vertical line: last point mismatch: got %v, want %v", got[len(got)-1], centroids[len(centroids)-1])
	}
}

func TestCompressCentroids_BoundaryBucketing_NoDownsample(t *testing.T) {
	// Construct a box with width=2, height=2 and choose maxPoints so no downsampling occurs.
	// cellSize = sqrt((2*2)/maxPoints). For maxPoints=4, cellSize=1. To avoid downsampling,
	// we pick maxPoints larger than the number of unique cells we insert.
	maxPoints := 10
	pts := [][2]float64{
		{0, 0}, {0, 0}, // duplicate in same cell
		{1, 0}, {1, 0}, // duplicate on edge cell
		{0, 1}, {1, 1},
		{2, 0}, {0, 2}, {2, 2}, {1, 2}, {2, 1},
	}

	got := CompressCentroids(pts, maxPoints)

	// Build set of unique input cells using known cell size 1 (see comment). We can directly
	// check unique coordinates here because all chosen points are on integer grid cells.
	seen := make(map[[2]float64]int)
	for _, p := range got {
		seen[p]++
	}
	// We expect duplicates removed; count of unique outputs should be number of unique inputs on grid.
	wantUnique := map[[2]float64]struct{}{
		{0, 0}: {}, {1, 0}: {}, {0, 1}: {}, {1, 1}: {}, {2, 0}: {}, {0, 2}: {}, {2, 2}: {}, {1, 2}: {}, {2, 1}: {},
	}
	if len(seen) != len(wantUnique) {
		t.Fatalf("boundary bucketing: got %d unique, want %d; got=%v", len(seen), len(wantUnique), got)
	}
	for p := range wantUnique {
		if seen[p] != 1 {
			t.Fatalf("boundary bucketing: expected point %v exactly once, got count=%d", p, seen[p])
		}
	}
}

func TestCompressCentroids_MaxPointsOne(t *testing.T) {
	centroids := [][2]float64{{0, 0}, {1, 1}, {2, 2}, {3, 3}}
	got := CompressCentroids(centroids, 1)
	if len(got) != 1 {
		t.Fatalf("maxPoints=1: got len=%d, want 1; got=%v", len(got), got)
	}
	// Expect first input representative preserved
	if got[0] != centroids[0] {
		t.Fatalf("maxPoints=1: first point mismatch: got %v, want %v", got[0], centroids[0])
	}
}

func TestCompressCentroids_MaxPointsTwo(t *testing.T) {
	centroids := [][2]float64{{0, 0}, {1, 1}, {2, 2}, {700, 500}, {909, 1560}, {2, 2}, {1, 1}, {0, 0}}
	got := CompressCentroids(centroids, 3)
	if len(got) != 3 {
		t.Fatalf("maxPoints=3: got len=%d, want 3; got=%v", len(got), got)
	}
	// Expect first input representative preserved
	if got[0] != centroids[0] {
		t.Fatalf("maxPoints=3: first point mismatch: got %v, want %v", got[0], centroids[0])
	}
}

func TestCompressCentroids_MaxPointsEqualToLength(t *testing.T) {
	centroids := [][2]float64{{0, 0}, {1, 1}, {2, 2}, {700, 500}, {909, 1560}, {2, 2}, {1, 1}, {0, 0}}
	got := CompressCentroids(centroids, 8)
	if len(got) != 8 {
		t.Fatalf("maxPoints=8: got len=%d, want 8; got=%v", len(got), got)
	}
	// Expect first input representative preserved
	if got[0] != centroids[0] {
		t.Fatalf("maxPoints=8: first point mismatch: got %v, want %v", got[0], centroids[0])
	}
}

// Testing if points equally spaced returns expected count when maxPoints does not evenly divide.
func TestCompressCentroids_MaxPointsDividedEqually(t *testing.T) {
	centroids := [][2]float64{{0, 0}, {0, 1}, {1, 0}, {2, 0}, {0, 2}, {2, 2}, {1, 1}, {1, 2}}
	got := CompressCentroids(centroids, 4)
	if len(got) != 4 {
		t.Fatalf("maxPoints=4: got len=%d, want 4; got=%v", len(got), got)
	}

	// Expect first input representative preserved
	if got[0] != centroids[0] {
		t.Fatalf("maxPoints=4: first point mismatch: got %v, want %v", got[0], centroids[0])
	}
}

// Explicitly exercise the zero-area fallback which returns the first point.
// This mirrors the guard that handles `width==0 && height==0` (and the inner
// `maxRange<=0` path) to ensure we always return exactly one representative.
func TestCompressCentroids_ZeroAreaFallback(t *testing.T) {
	pts := [][2]float64{{42, 42}, {42, 42}, {42, 42}, {42, 42}}
	// Use maxPoints < len(pts) so we actually invoke the zero-area fallback
	// (the function returns input as-is when len(pts) <= maxPoints).
	got := CompressCentroids(pts, 2)
	if len(got) != 1 {
		t.Fatalf("zero-area fallback: got len=%d, want 1; got=%v", len(got), got)
	}
	if got[0] != pts[0] {
		t.Fatalf("zero-area fallback: first point mismatch: got %v, want %v", got[0], pts[0])
	}
}

// helper to recompute the intermediate reduced set like the implementation, for property checks.
func recomputeReduced(centroids [][2]float64, maxPoints int) (reduced [][2]float64) {
	if len(centroids) == 0 {
		return nil
	}
	minX, maxX := centroids[0][0], centroids[0][0]
	minY, maxY := centroids[0][1], centroids[0][1]
	for _, c := range centroids[1:] {
		if c[0] < minX {
			minX = c[0]
		}
		if c[0] > maxX {
			maxX = c[0]
		}
		if c[1] < minY {
			minY = c[1]
		}
		if c[1] > maxY {
			maxY = c[1]
		}
	}
	width := maxX - minX
	height := maxY - minY
	if width == 0 && height == 0 {
		return centroids[:1]
	}
	cellSize := math.Sqrt((width * height) / float64(maxPoints))
	if cellSize <= 0 {
		maxRange := math.Max(width, height)
		if maxRange <= 0 {
			return centroids[:1]
		}
		cellSize = maxRange / float64(maxPoints)
	}
	type cellKey struct{ x, y int }
	seen := map[cellKey]struct{}{}
	for _, c := range centroids {
		key := cellKey{
			x: int(math.Floor((c[0] - minX) / cellSize)),
			y: int(math.Floor((c[1] - minY) / cellSize)),
		}
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		reduced = append(reduced, c)
	}
	return reduced
}

func TestCompressCentroids_Property_Randomized(t *testing.T) {
	r := mrand.New(mrand.NewSource(42))
	cases := 10
	for ci := 0; ci < cases; ci++ {
		n := 200
		// random ranges including negatives and large values
		baseX := (r.Float64()*2 - 1) * 1e6 // [-1e6, 1e6]
		baseY := (r.Float64()*2 - 1) * 1e6
		spanX := r.Float64()*1e6 + 1 // [1, 1e6+1]
		spanY := r.Float64()*1e6 + 1
		centroids := make([][2]float64, n)
		for i := 0; i < n; i++ {
			cx := baseX + r.Float64()*spanX
			cy := baseY + r.Float64()*spanY
			centroids[i] = [2]float64{cx, cy}
		}
		maxCandidates := []int{1, 5, 20, 50, n, n + 10}
		for _, maxPoints := range maxCandidates {
			got := CompressCentroids(centroids, maxPoints)

			// Invariants
			if maxPoints <= 0 || len(centroids) <= maxPoints {
				if !reflect.DeepEqual(got, centroids) {
					t.Fatalf("prop case: expected passthrough for maxPoints=%d; got len=%d, want len=%d", maxPoints, len(got), len(centroids))
				}
				continue
			}
			if len(got) > maxPoints {
				t.Fatalf("prop case: len(got)=%d exceeds maxPoints=%d", len(got), maxPoints)
			}
			// got subset of input
			inSet := make(map[[2]float64]struct{}, len(centroids))
			for _, p := range centroids {
				inSet[p] = struct{}{}
			}
			for _, p := range got {
				if _, ok := inSet[p]; !ok {
					t.Fatalf("prop case: output point %v not in input", p)
				}
			}

			// If downsampling occurred (reduced > maxPoints), endpoints should match reduced endpoints
			reduced := recomputeReduced(centroids, maxPoints)
			if len(reduced) > maxPoints {
				if got[0] != reduced[0] {
					t.Fatalf("prop case: first point not preserved: got %v, want %v", got[0], reduced[0])
				}
				if maxPoints > 1 {
					// For maxPoints==1, only the first element can be present.
					last := reduced[len(reduced)-1]
					foundLast := false
					for _, p := range got {
						if p == last {
							foundLast = true
							break
						}
					}
					if !foundLast {
						t.Fatalf("prop case: expected reduced last point %v to appear in output", last)
					}
				}
			}
		}
	}
}

// Ensure when downsampling occurs, the last element of the output corresponds
// exactly to the last element of the reduced set (after bucketing). This
// specifically validates the index clamping logic used when rounding could
// overshoot the last index.
func TestCompressCentroids_Downsample_LastIndexClamped(t *testing.T) {
	// Build a dense set to force reduced > maxPoints via bucketing.
	n := 500
	pts := make([][2]float64, 0, n)
	for i := 0; i < n; i++ {
		// Spread points across a rectangle; ensure duplicates to reduce via cells
		x := float64(i % 50) // 0..49
		y := float64(i % 40) // 0..39
		pts = append(pts, [2]float64{x, y})
	}

	maxPoints := 25
	got := CompressCentroids(pts, maxPoints)

	// Recompute reduced as in implementation to inspect endpoints.
	reduced := recomputeReduced(pts, maxPoints)
	if len(reduced) <= maxPoints {
		t.Fatalf("setup failed: expected reduced > maxPoints; got reduced=%d, maxPoints=%d", len(reduced), maxPoints)
	}

	// First point should match reduced[0]
	if got[0] != reduced[0] {
		t.Fatalf("last-index clamp: first point mismatch: got %v, want %v", got[0], reduced[0])
	}
	// Last point should be exactly the reduced last element. This covers the
	// clamping behavior in case rounding overshoots the last index.
	if got[len(got)-1] != reduced[len(reduced)-1] {
		t.Fatalf("last-index clamp: last point mismatch: got %v, want %v", got[len(got)-1], reduced[len(reduced)-1])
	}
}

// --- BuildCentroids tests ---

func almostEq(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func assertPointsAlmostEqual(t *testing.T, got, want [][2]float64, eps float64) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("len mismatch: got %d, want %d; got=%v want=%v", len(got), len(want), got, want)
	}
	for i := range got {
		if !almostEq(got[i][0], want[i][0], eps) || !almostEq(got[i][1], want[i][1], eps) {
			t.Fatalf("point %d mismatch: got %v, want %v", i, got[i], want[i])
		}
	}
}

func TestBuildCentroids_NoNormalization(t *testing.T) {
	// Two boxes: centers are (5,5) and (20,30)
	traject := []interface{}{
		[]interface{}{0.0, 0.0, 10.0, 10.0},
		[]interface{}{10.0, 20.0, 30.0, 40.0},
	}
	got := BuildCentroids(traject, 0, 0)
	want := [][2]float64{{5, 5}, {20, 30}}
	assertPointsAlmostEqual(t, got, want, 1e-9)
}

func TestBuildCentroids_Normalization(t *testing.T) {
	// Same boxes; frame 50x100 -> centers scaled to 100x100 space
	traject := []interface{}{
		[]interface{}{0.0, 0.0, 10.0, 10.0},   // center (5,5) -> (10,5)
		[]interface{}{10.0, 20.0, 30.0, 40.0}, // center (20,30) -> (40,30)
	}
	got := BuildCentroids(traject, 50, 100)
	want := [][2]float64{{10, 5}, {40, 30}}
	assertPointsAlmostEqual(t, got, want, 1e-9)
}

func TestBuildCentroids_SkipInvalidEntries(t *testing.T) {
	traject := []interface{}{
		123,                                 // not []interface{}
		[]interface{}{1.0, 2.0, 3.0},        // too short
		[]interface{}{0.0, 0.0, 10.0, 10.0}, // valid -> center (5,5)
	}
	got := BuildCentroids(traject, 0, 0)
	want := [][2]float64{{5, 5}}
	assertPointsAlmostEqual(t, got, want, 1e-9)
}

func TestBuildCentroids_OneDimensionProvided(t *testing.T) {
	// frameWidth>0 but frameHeight=0 -> no normalization should happen per implementation
	traject := []interface{}{
		[]interface{}{0.0, 0.0, 10.0, 10.0}, // center (5,5)
	}
	got := BuildCentroids(traject, 50, 0)
	want := [][2]float64{{5, 5}}
	assertPointsAlmostEqual(t, got, want, 1e-9)
}
