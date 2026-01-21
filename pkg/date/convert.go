package date

import (
	"fmt"
	"strings"
	"time"
)

func GetHour(timezone string, timestamp int64) int {
	t := time.Unix(timestamp, 0)
	loc, _ := time.LoadLocation(timezone)
	return t.In(loc).Hour()
}

func GetDate(timezone string, timestamp int64) string {
	t := time.Unix(timestamp, 0)
	loc, _ := time.LoadLocation(timezone)
	return t.In(loc).Format("02-01-2006")
}

func GetTime(timezone string, timestamp int64) string {
	t := time.Unix(timestamp, 0)
	loc, _ := time.LoadLocation(timezone)
	return t.In(loc).Format("15:04:05")
}

func GetDateTime(timezone string, timestamp int64) string {
	t := time.Unix(timestamp, 0)
	loc, _ := time.LoadLocation(timezone)
	return t.In(loc).Format("02-01-2006 - 15:04:05")
}

func GetDateTimeLong(timezone string, timestamp int64) string {
	t := time.Unix(timestamp, 0)
	loc, _ := time.LoadLocation(timezone)
	timeInLocation := t.In(loc)

	suffix := "th"
	switch timeInLocation.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}

	return timeInLocation.Format("January 2" + suffix + " 2006, 15:04:05")
}

func GetDateShort(timezone string, timestamp int64) string {
	t := time.Unix(timestamp, 0)
	loc, _ := time.LoadLocation(timezone)
	timeInLocation := t.In(loc)

	suffix := "th"
	switch timeInLocation.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}

	return timeInLocation.Format("January 2" + suffix + ",Monday")
}

func GetTimestamp(timezone string, date string) int64 {
	layout := "02-01-2006"
	loc, _ := time.LoadLocation(timezone)
	t, err := time.ParseInLocation(layout, date, loc)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return t.Unix()
}

// FormatDuration formats a float64 duration (in seconds) into hh:mm:ss format if hours are greater than 0,
// otherwise it returns mm:ss format.
func FormatDuration(duration float64) string {
	if duration == 0 {
		return "00:00"
	}
	totalSeconds := int(duration) // Convert float64 to int
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds) // Include hours if greater than 0
	}
	if minutes == 0 && seconds == 0 {
		return "<00:01"
	}
	return fmt.Sprintf("%02d:%02d", minutes, seconds) // Omit hours if 0
}

// FormatDurationShortMillis formats a duration in milliseconds into a short, human-readable string
// like "30s" or "4m 12s".
func FormatDurationShortMillis(durationMs int) string {
	if durationMs <= 0 {
		return "0s"
	}
	totalSeconds := durationMs / 1000
	if totalSeconds == 0 {
		return "<1s"
	}
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	var parts []string
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%ds", seconds))
	}
	if len(parts) == 0 {
		return "0s"
	}
	return strings.Join(parts, " ")
}
