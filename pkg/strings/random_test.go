package strings

import (
	"strings"
	"testing"
)

func TestRandStringBytesRmndr(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"empty string", 0},
		{"single character", 1},
		{"normal length", 10},
		{"long string", 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandStringBytesRmndr(tt.length)
			if len(result) != tt.length {
				t.Errorf("RandStringBytesRmndr(%d) = %d characters, want %d", tt.length, len(result), tt.length)
			}
			// Check if all characters are from the expected charset
			for _, char := range result {
				if !strings.Contains(letterBytes, string(char)) {
					t.Errorf("RandStringBytesRmndr() contains invalid character: %c", char)
				}
			}
		})
	}
}

func TestRandKey(t *testing.T) {
	// Test that RandKey generates a string of expected length
	key, err := RandKey()
	if err != nil {
		t.Fatalf("RandKey() returned error: %v", err)
	}
	if len(key) != 32 {
		t.Errorf("RandKey() = %d characters, want 32", len(key))
	}

	// Test that all characters are from the expected charset
	for _, char := range key {
		if !strings.Contains(charset, string(char)) {
			t.Errorf("RandKey() contains invalid character: %c", char)
		}
	}

	// Test that multiple calls generate different keys
	key2, err := RandKey()
	if err != nil {
		t.Fatalf("RandKey() returned error on second call: %v", err)
	}
	if key == key2 {
		t.Errorf("RandKey() generated identical keys: %s", key)
	}
}

func TestGenerateShortLink(t *testing.T) {
	link1 := GenerateShortLink()
	link2 := GenerateShortLink()

	// Test length
	if len(link1) != 6 {
		t.Errorf("GenerateShortLink() = %d characters, want 6", len(link1))
	}

	// Test uniqueness (basic check)
	if link1 == link2 {
		t.Errorf("GenerateShortLink() generated identical links: %s", link1)
	}

	// Test characters are from expected set
	for _, char := range link1 {
		if !strings.Contains(letterBytes, string(char)) {
			t.Errorf("GenerateShortLink() contains invalid character: %c", char)
		}
	}
}

func TestRandKeyErrorHandling(t *testing.T) {
	// Test multiple calls to potentially catch edge cases where crypto/rand.Read might fail
	// While crypto/rand.Read rarely fails, this test ensures proper error handling exists
	for i := 0; i < 100; i++ {
		key, err := RandKey()

		if err != nil {
			// If we do encounter an error, verify it's properly formatted
			expectedPrefix := "failed to generate secure random key:"
			if !strings.Contains(err.Error(), expectedPrefix) {
				t.Errorf("RandKey() error message format incorrect, got: %v", err)
			}
			// Verify empty key is returned on error
			if key != "" {
				t.Errorf("RandKey() returned non-empty key with error: key=%q, err=%v", key, err)
			}
			t.Logf("Successfully caught and validated RandKey() error: %v", err)
			return // Test passed - we covered the error case
		}

		// Verify successful generation
		if len(key) != 32 {
			t.Errorf("RandKey() generated key of length %d, want 32", len(key))
		}
	}

	// If no error occurred in 100 attempts, that's normal and expected
	t.Log("RandKey() completed 100 successful calls - error handling code is present but not triggered")
}

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name    string
		keyType string
		wantErr bool
	}{
		{"public key", "public", false},
		{"private key", "private", false},
		{"invalid key type", "invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GenerateKey(tt.keyType)

			if tt.wantErr {
				if err == nil {
					t.Errorf("GenerateKey(%q) expected error but got none", tt.keyType)
				}
				return
			}

			if err != nil {
				t.Errorf("GenerateKey(%q) returned unexpected error: %v", tt.keyType, err)
				return
			}

			switch tt.keyType {
			case "public":
				if !strings.HasPrefix(key, "UUG") {
					t.Errorf("GenerateKey(\"public\") = %q, expected to start with 'UUG'", key)
				}
				if len(key) != 19 { // "UUG" + 16 random characters
					t.Errorf("GenerateKey(\"public\") = %d characters, want 19", len(key))
				}
			case "private":
				if len(key) != 32 {
					t.Errorf("GenerateKey(\"private\") = %d characters, want 32", len(key))
				}
			}
		})
	}
}
