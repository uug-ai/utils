package main

import (
	"bytes"
	"io"
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

func TestMain_NotSet(t *testing.T) {
	_ = os.Unsetenv("MY_ENV_VAR")

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w
	defer func() {
		_ = w.Close()
		os.Stdout = old
		_ = r.Close()
	}()

	main()

	_ = w.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}
	got := buf.String()
	want := "Environment variable MY_ENV_VAR is not set.\n"
	if got != want {
		t.Fatalf("unexpected output: got %q, want %q", got, want)
	}
}

func TestMain_Set(t *testing.T) {
	t.Setenv("MY_ENV_VAR", "from-main")

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w
	defer func() {
		_ = w.Close()
		os.Stdout = old
		_ = r.Close()
	}()

	main()

	_ = w.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}
	got := buf.String()
	want := "The value of MY_ENV_VAR is: from-main\n"
	if got != want {
		t.Fatalf("unexpected output: got %q, want %q", got, want)
	}
}
