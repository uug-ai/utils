package ustrings

import "testing"

func TestCapitalize(t *testing.T) {
	if Capitalize("hello") != "Hello" {
		t.Errorf("Capitalize did not uppercase first rune")
	}
	if Capitalize("") != "" {
		t.Errorf("Capitalize should return empty string unchanged")
	}
	if Capitalize("hELLO") != "HELLO" {
		t.Errorf("Capitalize should only change first rune")
	}
}

func TestToCamelCase(t *testing.T) {
	cases := map[string]string{
		"hello_world": "helloWorld",
		"Hello-world": "helloWorld",
		"hello world": "helloWorld",
		"HTTPServer":  "httpServer",
	}
	for input, expected := range cases {
		if got := ToCamelCase(input); got != expected {
			t.Errorf("ToCamelCase(%q) = %q, want %q", input, got, expected)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	cases := map[string]string{
		"HelloWorld":   "hello_world",
		"hello-world":  "hello_world",
		"HTTPServer":   "http_server",
		"already_snake": "already_snake",
	}
	for input, expected := range cases {
		if got := ToSnakeCase(input); got != expected {
			t.Errorf("ToSnakeCase(%q) = %q, want %q", input, got, expected)
		}
	}
}

func TestToKebabCase(t *testing.T) {
	cases := map[string]string{
		"HelloWorld":  "hello-world",
		"hello_world": "hello-world",
	}
	for input, expected := range cases {
		if got := ToKebabCase(input); got != expected {
			t.Errorf("ToKebabCase(%q) = %q, want %q", input, got, expected)
		}
	}
}

func TestPadLeftRight(t *testing.T) {
	if PadLeft("42", "0", 5) != "00042" {
		t.Errorf("PadLeft did not pad correctly")
	}
	if PadRight("42", "0", 5) != "42000" {
		t.Errorf("PadRight did not pad correctly")
	}
	if PadLeft("abc", "x", 2) != "abc" {
		t.Errorf("PadLeft should not pad when target shorter")
	}
	if PadLeft("abc", "", 5) != "abc" {
		t.Errorf("PadLeft should not pad with empty pad string")
	}
}

func TestTruncate(t *testing.T) {
	if Truncate("hello", 10, "...") != "hello" {
		t.Errorf("Truncate should keep shorter string")
	}
	if Truncate("hello world", 5, "...") != "hello..." {
		t.Errorf("Truncate did not add ellipsis")
	}
	if Truncate("test", 0, "...") != "..." {
		t.Errorf("Truncate should return ellipsis when length is zero")
	}
	if Truncate("hello world", 5, " [more]") != "hello [more]" {
		t.Errorf("Truncate should allow custom ellipsis")
	}
}

func TestEllipsisMiddle(t *testing.T) {
	if EllipsisMiddle("hello world", 7) != "he...ld" {
		t.Errorf("EllipsisMiddle did not truncate as expected")
	}
	if EllipsisMiddle("short", 10) != "short" {
		t.Errorf("EllipsisMiddle should return unchanged if shorter")
	}
	if EllipsisMiddle("hello", 3) != "hel" {
		t.Errorf("EllipsisMiddle should return first runes when too short")
	}
}

func TestReverse(t *testing.T) {
	if Reverse("world!") != "!dlrow" {
		t.Errorf("Reverse did not reverse ascii string")
	}
	if Reverse("\U0001F44B\U0001F30D") != "\U0001F30D\U0001F44B" {
		t.Errorf("Reverse did not reverse utf-8 string")
	}
}

func TestRemoveWhitespace(t *testing.T) {
	if RemoveWhitespace(" a \t b\n") != "ab" {
		t.Errorf("RemoveWhitespace did not remove whitespace")
	}
}

func TestSplitAndTrim(t *testing.T) {
	parts := SplitAndTrim(" a, b ,, c ", ",")
	if len(parts) != 3 || parts[0] != "a" || parts[1] != "b" || parts[2] != "c" {
		t.Errorf("SplitAndTrim did not return expected parts")
	}
}

func TestJoinNonEmpty(t *testing.T) {
	joined := JoinNonEmpty([]string{"a", "", "b"}, ",")
	if joined != "a,b" {
		t.Errorf("JoinNonEmpty did not skip empty strings")
	}
}

func TestDefaultIfEmpty(t *testing.T) {
	if DefaultIfEmpty("", "x") != "x" {
		t.Errorf("DefaultIfEmpty did not return fallback")
	}
	if DefaultIfEmpty("y", "x") != "y" {
		t.Errorf("DefaultIfEmpty should return original value")
	}
}
