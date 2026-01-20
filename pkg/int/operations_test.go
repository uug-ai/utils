package int

import "testing"

func TestToInt(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		fallback int
		expected int
	}{
		{"int value", int(42), 0, 42},
		{"int32 value", int32(7), 0, 7},
		{"int64 value", int64(8), 0, 8},
		{"float32 truncates", float32(3.9), 0, 3},
		{"float64 truncates", float64(9.99), 0, 9},
		{"string falls back", "5", -1, -1},
		{"bool falls back", true, 13, 13},
		{"nil falls back", nil, 99, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToInt(tt.value, tt.fallback)
			if result != tt.expected {
				t.Errorf("ToInt(%v, %d) = %d, want %d", tt.value, tt.fallback, result, tt.expected)
			}
		})
	}
}
