package geometry

import "math"

func CompressCentroids(centroids [][2]float64, maxPoints int) [][2]float64 {
	if maxPoints <= 0 || len(centroids) <= maxPoints {
		return centroids
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

	type cellKey struct {
		x int
		y int
	}
	seen := make(map[cellKey]struct{}, maxPoints)
	reduced := make([][2]float64, 0, maxPoints)
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

	if len(reduced) <= maxPoints {
		return reduced
	}

	step := float64(len(reduced)-1) / float64(maxPoints-1)
	downsampled := make([][2]float64, 0, maxPoints)
	for i := 0; i < maxPoints; i++ {
		idx := int(math.Round(float64(i) * step))
		if idx >= len(reduced) {
			idx = len(reduced) - 1
		}
		downsampled = append(downsampled, reduced[idx])
	}
	return downsampled
}
