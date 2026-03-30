package ustrings

import (
	"strings"
	"unicode"
)

// Capitalize uppercases the first letter of the string.
func Capitalize(value string) string {
	if value == "" {
		return value
	}
	runes := []rune(value)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// ToCamelCase converts a string to lower camel case.
func ToCamelCase(value string) string {
	words := splitWords(value)
	if len(words) == 0 {
		return ""
	}
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	words[0] = strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		words[i] = Capitalize(words[i])
	}
	return strings.Join(words, "")
}

// ToSnakeCase converts a string to snake_case.
func ToSnakeCase(value string) string {
	words := splitWords(value)
	if len(words) == 0 {
		return ""
	}
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "_")
}

// ToKebabCase converts a string to kebab-case.
func ToKebabCase(value string) string {
	words := splitWords(value)
	if len(words) == 0 {
		return ""
	}
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "-")
}

// PadLeft pads a string to a specific length using padWith on the left side.
func PadLeft(value, padWith string, padTo int) string {
	if padTo <= len(value) || padWith == "" {
		return value
	}
	padding := getPaddingString(padWith, padTo-len(value))
	return padding + value
}

// PadRight pads a string to a specific length using padWith on the right side.
func PadRight(value, padWith string, padTo int) string {
	if padTo <= len(value) || padWith == "" {
		return value
	}
	padding := getPaddingString(padWith, padTo-len(value))
	return value + padding
}

// Truncate truncates a string to maxLength and appends ellipsis if needed.
func Truncate(value string, maxLength int, ellipsis string) string {
	runes := []rune(value)
	if maxLength <= 0 {
		return ellipsis
	}
	if len(runes) <= maxLength {
		return value
	}
	if ellipsis == "" {
		return string(runes[:maxLength])
	}
	runeEllipsis := []rune(ellipsis)
	if len(runeEllipsis) > maxLength {
		return string(runeEllipsis[:maxLength])
	}
	return string(runes[:maxLength]) + ellipsis
}

// EllipsisMiddle truncates the middle of a string with "..." to maxLength.
func EllipsisMiddle(value string, maxLength int) string {
	runes := []rune(value)
	if maxLength <= 0 {
		return ""
	}
	if len(runes) <= maxLength {
		return value
	}
	if maxLength <= 3 {
		return string(runes[:maxLength])
	}
	leftCount := (maxLength - 3) / 2
	rightCount := maxLength - 3 - leftCount
	left := string(runes[:leftCount])
	right := string(runes[len(runes)-rightCount:])
	return left + "..." + right
}

// Reverse reverses a UTF-8 string.
func Reverse(value string) string {
	runes := []rune(value)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// RemoveWhitespace removes all whitespace characters from the string.
func RemoveWhitespace(value string) string {
	parts := strings.Fields(value)
	return strings.Join(parts, "")
}

// SplitAndTrim splits a string by sep and trims whitespace from each part.
func SplitAndTrim(value, sep string) []string {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, sep)
	output := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		output = append(output, trimmed)
	}
	return output
}

// JoinNonEmpty joins non-empty strings with the given separator.
func JoinNonEmpty(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	filtered := make([]string, 0, len(parts))
	for _, part := range parts {
		if part == "" {
			continue
		}
		filtered = append(filtered, part)
	}
	return strings.Join(filtered, sep)
}

// DefaultIfEmpty returns fallback if value is empty.
func DefaultIfEmpty(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

func getPaddingString(padWith string, padLength int) string {
	if padLength <= 0 || padWith == "" {
		return ""
	}
	builder := strings.Builder{}
	builder.Grow(padLength)
	for builder.Len() < padLength {
		builder.WriteString(padWith)
	}
	padding := builder.String()
	if len(padding) > padLength {
		return padding[:padLength]
	}
	return padding
}

func splitWords(value string) []string {
	runes := []rune(value)
	words := make([]string, 0, len(runes))
	var current []rune
	flush := func() {
		if len(current) == 0 {
			return
		}
		words = append(words, string(current))
		current = current[:0]
	}
	for i, r := range runes {
		if isSeparator(r) {
			flush()
			continue
		}
		if len(current) == 0 {
			current = append(current, r)
			continue
		}
		prev := runes[i-1]
		var next rune
		if i+1 < len(runes) {
			next = runes[i+1]
		}
		if isBoundary(prev, r, next) {
			flush()
		}
		current = append(current, r)
	}
	flush()
	return words
}

func isSeparator(r rune) bool {
	if unicode.IsSpace(r) {
		return true
	}
	switch r {
	case '_', '-', '.', '/', ':':
		return true
	default:
		return false
	}
}

func isBoundary(prev, current, next rune) bool {
	if unicode.IsLower(prev) && unicode.IsUpper(current) {
		return true
	}
	if unicode.IsUpper(prev) && unicode.IsUpper(current) && next != 0 && unicode.IsLower(next) {
		return true
	}
	if unicode.IsDigit(prev) && unicode.IsLetter(current) {
		return true
	}
	if unicode.IsLetter(prev) && unicode.IsDigit(current) {
		return true
	}
	return false
}
