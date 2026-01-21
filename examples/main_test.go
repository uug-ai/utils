package main

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintEnvVar_NotSet(t *testing.T) {
	// Ensure variable is unset
	_ = os.Unsetenv("MY_ENV_VAR")
	var buf bytes.Buffer

	printEnvVar("MY_ENV_VAR", &buf)

	got := buf.String()
	want := "Environment variable MY_ENV_VAR is not set.\n"
	if got != want {
		t.Fatalf("unexpected output: got %q, want %q", got, want)
	}
}

func TestPrintEnvVar_Set(t *testing.T) {
	// Set variable
	t.Setenv("MY_ENV_VAR", "hello-world")
	var buf bytes.Buffer

	printEnvVar("MY_ENV_VAR", &buf)

	got := buf.String()
	want := "The value of MY_ENV_VAR is: hello-world\n"
	if got != want {
		t.Fatalf("unexpected output: got %q, want %q", got, want)
	}
}
