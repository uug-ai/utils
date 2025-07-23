package array

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		target   string
		expected bool
	}{
		{"empty slice", []string{}, "test", false},
		{"contains string", []string{"apple", "banana", "cherry"}, "banana", true},
		{"does not contain", []string{"apple", "banana", "cherry"}, "grape", false},
		{"empty string in slice", []string{"apple", "", "cherry"}, "", true},
		{"case sensitive", []string{"Apple", "Banana"}, "apple", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.slice, tt.target)
			if result != tt.expected {
				t.Errorf("Contains(%v, %q) = %v, want %v", tt.slice, tt.target, result, tt.expected)
			}
		})
	}
}

func TestUniq(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"no duplicates", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"with duplicates", []string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}},
		{"empty slice", []string{}, []string{}},
		{"all same", []string{"a", "a", "a"}, []string{"a"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uniq(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("Uniq(%v) returned %d items, want %d", tt.input, len(result), len(tt.expected))
				return
			}

			// Convert to maps for comparison since order may vary
			resultMap := make(map[string]bool)
			expectedMap := make(map[string]bool)

			for _, v := range result {
				resultMap[v] = true
			}
			for _, v := range tt.expected {
				expectedMap[v] = true
			}

			for k := range expectedMap {
				if !resultMap[k] {
					t.Errorf("Uniq(%v) missing expected item: %s", tt.input, k)
				}
			}

			for k := range resultMap {
				if !expectedMap[k] {
					t.Errorf("Uniq(%v) contains unexpected item: %s", tt.input, k)
				}
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			"basic difference",
			[]string{"a", "b", "c"},
			[]string{"b", "d"},
			[]string{"a", "c"},
		},
		{
			"no difference",
			[]string{"a", "b"},
			[]string{"a", "b", "c"},
			[]string{},
		},
		{
			"empty second slice",
			[]string{"a", "b"},
			[]string{},
			[]string{"a", "b"},
		},
		{
			"empty first slice",
			[]string{},
			[]string{"a", "b"},
			[]string{},
		},
		{
			"both empty",
			[]string{},
			[]string{},
			[]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.slice1, tt.slice2)

			if len(result) != len(tt.expected) {
				t.Errorf("Difference(%v, %v) = %v, want %v", tt.slice1, tt.slice2, result, tt.expected)
				return
			}

			// Convert to maps for comparison
			resultMap := make(map[string]bool)
			expectedMap := make(map[string]bool)

			for _, v := range result {
				resultMap[v] = true
			}
			for _, v := range tt.expected {
				expectedMap[v] = true
			}

			for k := range expectedMap {
				if !resultMap[k] {
					t.Errorf("Difference(%v, %v) missing expected item: %s", tt.slice1, tt.slice2, k)
				}
			}

			for k := range resultMap {
				if !expectedMap[k] {
					t.Errorf("Difference(%v, %v) contains unexpected item: %s", tt.slice1, tt.slice2, k)
				}
			}
		})
	}
}
