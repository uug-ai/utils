package array

import "testing"

func TestArrayContainsAll(t *testing.T) {
	tests := []struct {
		name     string
		a        []string
		b        []string
		expected bool
	}{
		{"empty a slice", []string{}, []string{"x"}, true},
		{"all present", []string{"apple", "banana"}, []string{"banana", "cherry", "apple"}, true},
		{"some missing", []string{"apple", "grape"}, []string{"banana", "cherry", "apple"}, false},
		{"none present", []string{"kiwi"}, []string{"banana", "cherry", "apple"}, false},
		{"duplicates in a", []string{"apple", "apple"}, []string{"apple"}, true},
		{"duplicates in b", []string{"apple", "banana"}, []string{"apple", "apple", "banana"}, true},
		{"case sensitive", []string{"Apple"}, []string{"apple"}, false},
		{"empty b with non-empty a", []string{"a"}, []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ArrayContainsAll(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("ArrayContainsAll(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
