package string

import (
	"testing"
)

func TestToLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"all uppercase", "HELLO", "hello"},
		{"mixed case", "HeLLo WoRLd", "hello world"},
		{"all lowercase", "hello", "hello"},
		{"empty string", "", ""},
		{"numbers and symbols", "Hello123!@#", "hello123!@#"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToLower(tt.input)
			if result != tt.expected {
				t.Errorf("ToLower(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"positive number", "123", 123},
		{"negative number", "-456", -456},
		{"zero", "0", 0},
		{"invalid string", "abc", 0},
		{"empty string", "", 0},
		{"mixed alphanumeric", "123abc", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToInt(tt.input)
			if result != tt.expected {
				t.Errorf("StringToInt(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRemoveOrdinalSuffix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"1st", "January 1st 2023", "January 1 2023"},
		{"2nd", "March 22nd 2023", "March 22 2023"},
		{"3rd", "April 3rd 2023", "April 3 2023"},
		{"th suffix", "May 15th 2023", "May 15 2023"},
		{"multiple occurrences", "1st of January, 22nd of March", "1 of January, 22 of March"},
		{"no suffix", "January 15 2023", "January 15 2023"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveOrdinalSuffix(tt.input)
			if result != tt.expected {
				t.Errorf("RemoveOrdinalSuffix(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
