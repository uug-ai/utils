package date

import (
	"testing"
)

func TestGetDate(t *testing.T) {
	// Use a known timestamp: 2023-07-15 12:00:45 UTC (1689422445)
	timestamp := int64(1689422445)

	tests := []struct {
		name     string
		timezone string
		expected string
	}{
		{"UTC", "UTC", "15-07-2023"},
		{"New York", "America/New_York", "15-07-2023"}, // Should be same date
		{"London", "Europe/London", "15-07-2023"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDate(tt.timezone, timestamp)
			if result != tt.expected {
				t.Errorf("GetDate(%q, %d) = %q, want %q", tt.timezone, timestamp, result, tt.expected)
			}
		})
	}
}

func TestGetHour(t *testing.T) {
	// Use a known timestamp: 2023-07-15 12:00:45 UTC (1689422445)
	timestamp := int64(1689422445)

	tests := []struct {
		name     string
		timezone string
		expected int
	}{
		{"UTC", "UTC", 12},
		{"New York", "America/New_York", 8}, // UTC-4 in summer
		{"Tokyo", "Asia/Tokyo", 21},         // UTC+9
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetHour(tt.timezone, timestamp)
			if result != tt.expected {
				t.Errorf("GetHour(%q, %d) = %d, want %d", tt.timezone, timestamp, result, tt.expected)
			}
		})
	}
}

func TestGetTime(t *testing.T) {
	// Use a known timestamp: 2023-07-15 12:00:45 UTC (1689422445)
	timestamp := int64(1689422445)

	tests := []struct {
		name     string
		timezone string
		expected string
	}{
		{"UTC", "UTC", "12:00:45"},
		{"New York", "America/New_York", "08:00:45"}, // UTC-4 in summer
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetTime(tt.timezone, timestamp)
			if result != tt.expected {
				t.Errorf("GetTime(%q, %d) = %q, want %q", tt.timezone, timestamp, result, tt.expected)
			}
		})
	}
}

func TestGetDateTime(t *testing.T) {
	// Use a known timestamp: 2023-07-15 12:00:45 UTC (1689422445)
	timestamp := int64(1689422445)

	tests := []struct {
		name     string
		timezone string
		expected string
	}{
		{"UTC", "UTC", "15-07-2023 - 12:00:45"},
		{"New York", "America/New_York", "15-07-2023 - 08:00:45"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDateTime(tt.timezone, timestamp)
			if result != tt.expected {
				t.Errorf("GetDateTime(%q, %d) = %q, want %q", tt.timezone, timestamp, result, tt.expected)
			}
		})
	}
}

func TestGetDateTimeLong(t *testing.T) {
	tests := []struct {
		name      string
		timezone  string
		timestamp int64
		expected  string
	}{
		{"1st day", "UTC", 1688126400, "June 30th 2023, 12:00:00"},  // 2023-06-30 12:00:00 UTC
		{"2nd day", "UTC", 1688212800, "July 1st 2023, 12:00:00"},   // 2023-07-01 12:00:00 UTC
		{"3rd day", "UTC", 1688299200, "July 2nd 2023, 12:00:00"},   // 2023-07-02 12:00:00 UTC
		{"21st day", "UTC", 1690027200, "July 22nd 2023, 12:00:00"}, // 2023-07-22 12:00:00 UTC
		{"22nd day", "UTC", 1690113600, "July 23rd 2023, 12:00:00"}, // 2023-07-23 12:00:00 UTC
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDateTimeLong(tt.timezone, tt.timestamp)
			if result != tt.expected {
				t.Errorf("GetDateTimeLong(%q, %d) = %q, want %q", tt.timezone, tt.timestamp, result, tt.expected)
			}
		})
	}
}

func TestGetDateShort(t *testing.T) {
	// Use a known timestamp: 2023-07-15 (Saturday) 12:00:45 UTC (1689422445)
	timestamp := int64(1689422445)

	result := GetDateShort("UTC", timestamp)
	expected := "July 15th,Saturday"

	if result != expected {
		t.Errorf("GetDateShort(\"UTC\", %d) = %q, want %q", timestamp, result, expected)
	}
}

func TestGetTimestamp(t *testing.T) {
	tests := []struct {
		name     string
		timezone string
		date     string
		expected int64
	}{
		{"UTC date", "UTC", "15-07-2023", 1689379200}, // 2023-07-15 00:00:00 UTC
		{"invalid date", "UTC", "invalid", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetTimestamp(tt.timezone, tt.date)
			if result != tt.expected {
				t.Errorf("GetTimestamp(%q, %q) = %d, want %d", tt.timezone, tt.date, result, tt.expected)
			}
		})
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration float64
		expected string
	}{
		{"zero duration", 0, "00:00"},
		{"less than one second", 0.5, "<00:01"},
		{"exactly one second", 1, "00:01"},
		{"one minute", 60, "01:00"},
		{"minutes and seconds", 125, "02:05"},
		{"one hour", 3600, "01:00:00"},
		{"hours, minutes, seconds", 3665, "01:01:05"},
		{"long duration", 7323, "02:02:03"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDuration(tt.duration)
			if result != tt.expected {
				t.Errorf("FormatDuration(%f) = %q, want %q", tt.duration, result, tt.expected)
			}
		})
	}
}
