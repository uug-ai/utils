package strings

import (
	"regexp"
	"strconv"
	"strings"
)

// ToLower converts a string to lowercase.
func ToLower(value string) string {
	lowerCase := strings.ToLower(value)
	return lowerCase
}

// StringToInt converts a string to an integer, returning 0 if conversion fails.
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// RemoveOrdinalSuffix removes ordinal suffixes (st, nd, rd, th) from date strings.
func RemoveOrdinalSuffix(dateStr string) string {
	re := regexp.MustCompile(`(\d+)(st|nd|rd|th)`)
	return re.ReplaceAllString(dateStr, `$1`)
}

// ObscureToken returns a token with its middle removed and replaced by "...".
// It keeps the first 5 and the last 5 characters. If the token is 10 characters
// or shorter it is returned unchanged.
func ObscureToken(token string) string {
	r := []rune(token)
	if len(r) <= 10 {
		return token
	}
	left := string(r[:5])
	right := string(r[len(r)-5:])
	return left + "..." + right
}

// ToStringSlice attempts to convert an interface{} to a []string.
func ToStringSlice(value any) []string {
	switch cast := value.(type) {
	case []string:
		return cast
	case []interface{}:
		output := make([]string, 0, len(cast))
		for _, item := range cast {
			if str, ok := item.(string); ok {
				output = append(output, str)
			}
		}
		return output
	default:
		return nil
	}
}
