package times

import (
	"log"
	"time"
	_ "time/tzdata"
)

const (
	defaultFormat    = "2006-01-02 15:04:05.000000"
	formatWithSlash  = "2006/01/02 15:04:05.000000"
	reverseWithSlash = "01/02/2006 15:04:05.000000"
)
const (
	dataWithSlashShort = "02/01/06"
	dataWithSlash      = "02/01/2006"
)

var utc6 = time.FixedZone("UTC+6", 6*60*60)

// TimeStampToAlmatyZone converts a timestamp to Almaty time zone
// If not exist Asia/Almaty time zone, return custom UTC+6 time zone
// return time in RFC3339 format
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

	loc := getAlmatyLocale()

	//если нет тайм зоны то проставляется utc
	// при добавлении зоны, ти приводится по ней bef: time 14:00 utc, after: 14:00 +06
	if parse.Location() == nil || parse.Location() == time.UTC {
		//если нет тайм зоны, то проставляется utc
		// при добавлении зоны, приводится по ней bef: time 14:00 00, after: 14:00 +06

		timeWithTimeZone := parse.In(loc).Add(-time.Hour * 6)
		formatted = timeWithTimeZone.Format(time.RFC3339)
	} else {
		timeWithTimeZone := parse.In(loc)
		formatted = timeWithTimeZone.Format(time.RFC3339)
	}

	return formatted, nil
}

// getAlmatyLocale returns the time.Location for Almaty
// if not exist return custom UTC+6 locale
func getAlmatyLocale() *time.Location {
	loc, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		log.Printf("error load location: %v", err)
		return utc6
	}
	return loc
}

func CheckAlmatyLocale() error {
	_, err := time.LoadLocation("Asia/Almaty")
	return err
}
func DateToAlmatyTime(str string) (string, error) {
	var timeFormats = []string{dataWithSlash, dataWithSlashShort, defaultFormat, formatWithSlash, time.RFC3339, reverseWithSlash}

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
	loc := getAlmatyLocale()
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
