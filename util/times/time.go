package times

import (
	"time"
)

const (
	defaultFormat   = "2006-01-02 15:04:05.000000"
	formatWithSlash = "2006/01/02 15:04:05.000000"
)

func TimeStampToAlmatyZone(str string) (string, error) {
	var parse time.Time
	parse, err := time.Parse(defaultFormat, str)
	if err != nil {
		parse, err = time.Parse(formatWithSlash, str)
		if err != nil {
			parse, err = time.Parse(time.RFC3339, str)
			if err != nil {
				return "", err

			}
		}
	}

	timeZone := "Asia/Almaty" // Example: Almaty, which is in the +6 time zone

	// Load the location
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return "", err
	}

	formatted := ""
	if parse.Location() == time.UTC {
		//если нет тайм зоны то проставляется utc
		// при добавлении зоны, ти приводится по ней bef: time 14:00 utc, after: 20:00 utc
		timeWithTimeZone := parse.In(loc).Add(-time.Hour * 6)
		formatted = timeWithTimeZone.Format(time.RFC3339)
	} else {
		timeWithTimeZone := parse.In(loc)
		formatted = timeWithTimeZone.Format(time.RFC3339)
	}

	return formatted, nil
}
