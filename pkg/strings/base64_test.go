package strings

import (
	"encoding/base64"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "aGVsbG8="},
		{"empty string", "", ""},
		{"with spaces", "hello world", "aGVsbG8gd29ybGQ="},
		{"special characters", "hello@123", "aGVsbG9AMTIz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Base64Encode(tt.input)
			if result != tt.expected {
				t.Errorf("Base64Encode(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBase64Decode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "aGVsbG8=", "hello"},
		{"empty string", "", ""},
		{"with spaces", "aGVsbG8gd29ybGQ=", "hello world"},
		{"special characters", "aGVsbG9AMTIz", "hello@123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Base64Decode(tt.input)
			if err != nil {
				t.Errorf("Base64Decode(%q) returned error: %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("Base64Decode(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestEncodeURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"simple string", "hello"},
		{"with spaces", "hello world"},
		{"special characters", "hello@123!"},
		{"URL", "https://example.com/path?query=value"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := EncodeURL(tt.input)
			// Verify it can be decoded back
			decoded, err := DecodeURL(encoded)
			if err != nil {
				t.Errorf("DecodeURL failed for encoded string: %v", err)
			}
			if decoded != tt.input {
				t.Errorf("EncodeURL/DecodeURL roundtrip failed: got %q, want %q", decoded, tt.input)
			}
		})
	}
}

func TestDecodeURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", base64.RawURLEncoding.EncodeToString([]byte("hello")), "hello"},
		{"with spaces", base64.RawURLEncoding.EncodeToString([]byte("hello world")), "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DecodeURL(tt.input)
			if err != nil {
				t.Errorf("DecodeURL(%q) returned error: %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("DecodeURL(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
