

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