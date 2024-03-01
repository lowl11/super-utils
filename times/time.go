package times

import (
	"time"
	_ "time/tzdata"
)

const (
	defaultGolangFormat                = "2006-01-02 15:04:05.999999999 -0700 MST"
	defaultFormat                      = "2006-01-02 15:04:05.000000"
	defaultFormatWithOutMillisecond    = "2006-01-02 15:04:05"
	formatWithSlash                    = "2006/01/02 15:04:05.000000"
	formatWithSlashWithOutMillisecond  = "2006/01/02 15:04:05"
	reverseWithSlash                   = "01/02/2006 15:04:05.000000"
	reverseWithSlashWithOutMillisecond = "01/02/2006 15:04:05"
)
const (
	dataWithSlashShort = "02/01/06"
	dataWithSlash      = "02/01/2006"
)

var almatyLocale *time.Location
var almatyOffset *int

// InitTimeZone initializes the time zone for Almaty
// offset is the time zone offset in hours
func InitTimeZone(offset int) {
	almatyOffset = &offset
	almatyLocale = time.FixedZone("Asia/Almaty", offset*60*60)
}

// TimeStampToAlmatyZone converts a timestamp to Almaty time zone
// If not exist Asia/Almaty time zone, return custom UTC+6 time zone
// return time in RFC3339 format
func TimeStampToAlmatyZone(str string) (string, error) {
	var timeFormats = []string{defaultGolangFormat, defaultFormat, formatWithSlash, time.RFC3339, reverseWithSlash,
		defaultFormatWithOutMillisecond, formatWithSlashWithOutMillisecond, reverseWithSlashWithOutMillisecond}

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

	loc, offset, err := GetAlmatyLocale()
	if err != nil {
		return "", err
	}

	//если нет тайм зоны то проставляется utc
	// при добавлении зоны, ти приводится по ней bef: time 14:00 utc, after: 14:00 +06
	if parse.Location() == nil || parse.Location() == time.UTC {
		//если нет тайм зоны, то проставляется utc
		// при добавлении зоны, приводится по ней bef: time 14:00 00, after: 14:00 +06

		timeWithTimeZone := parse.In(loc).Add(-time.Hour * time.Duration(offset))
		formatted = timeWithTimeZone.Format(time.RFC3339)
	} else {
		timeWithTimeZone := parse.In(loc)
		formatted = timeWithTimeZone.Format(time.RFC3339)
	}

	return formatted, nil
}

// GetAlmatyLocale returns the time.Location for Almaty
func GetAlmatyLocale() (*time.Location, int, error) {
	if almatyLocale == nil || almatyOffset == nil {
		return nil, 0, NotHasAlmatyZoneOffset
	}
	return almatyLocale, *almatyOffset, nil
}

func DateToAlmatyTime(str string) (string, error) {
	var timeFormats = []string{dataWithSlash, dataWithSlashShort, defaultFormat,
		formatWithSlash, time.RFC3339, reverseWithSlash, defaultGolangFormat}

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

	loc, offset, err := GetAlmatyLocale()
	if err != nil {
		return "", err
	}

	if parse.Location() == time.UTC {
		//если нет тайм зоны то проставляется utc
		// при добавлении зоны, ти приводится по ней bef: time 14:00 utc, after: 20:00 utc
		timeWithTimeZone := parse.In(loc).Add(-time.Hour * time.Duration(offset))
		formatted = timeWithTimeZone.Format(time.RFC3339)
	} else {
		timeWithTimeZone := parse.In(loc)
		formatted = timeWithTimeZone.Format(time.RFC3339)
	}

	return formatted, nil
}
