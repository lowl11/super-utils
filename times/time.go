package times

import (
	"time"
)

const (
	defaultFormat    = "2006-01-02 15:04:05.000000"
	formatWithSlash  = "2006/01/02 15:04:05.000000"
	reverseWithSlash = "02/01/2006 15:04:05.000000"
)
const (
	dataWithSlashShort = "02/01/06"
	dataWithSlash      = "02/01/2006"
)

func TimeStampToAlmatyZone(str string) (string, error) {
	var timeFormats = []string{defaultFormat, formatWithSlash, time.RFC3339, reverseWithSlash}

	var parse time.Time
	var err error

	for _, format := range timeFormats {
		parse, err = time.Parse(format, str)
		if err == nil {
			break
		}
	}

	if err != nil {
		return "", err
	}

	formatted := ""

	// +6 time zone
	loc := time.FixedZone("UTC+6", 6*60*60)

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

func DateToAlmatyTime(str string) (string, error) {
	var timeFormats = []string{dataWithSlash, dataWithSlashShort}

	var parse time.Time
	var err error

	for _, format := range timeFormats {
		parse, err = time.Parse(format, str)
		if err == nil {
			break
		}
	}

	if err != nil {
		return "", err
	}

	formatted := ""
	// +6 time zone
	loc := time.FixedZone("UTC+6", 6*60*60)
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
