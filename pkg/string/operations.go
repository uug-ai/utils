package string

import (
	"regexp"
	"strconv"
	"strings"
)

func ToLower(value string) string {
	lowerCase := strings.ToLower(value)
	return lowerCase
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func RemoveOrdinalSuffix(dateStr string) string {
	re := regexp.MustCompile(`(\d+)(st|nd|rd|th)`)
	return re.ReplaceAllString(dateStr, `$1`)
}
