package gogenerics

import "time"

func TryToDateTime(s string) (time.Time, bool) {
	var layouts = []string{
		// standard
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		// Handy time stamps.
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		//
		"2006-01-02 15:04:05",
	}
	for _, lo := range layouts {
		if tm, err := time.Parse(lo, s); err == nil {
			return tm, true
		}
	}
	return time.Time{}, false
}

func TryToDateUS(s string) (time.Time, bool) {
	var layouts = []string{
		"January 2, 2006",
		"Jan 2, 2006",
		"01/02/06",
		"01/02/2006",
		"Jan-02-06",
	}
	for _, lo := range layouts {
		if tm, err := time.Parse(lo, s); err == nil {
			return tm, true
		}
	}
	return time.Time{}, false
}

func TryToDateUK(s string) (time.Time, bool) {
	var layouts = []string{
		"2 January, 2006",
		"2 Jan, 2006",
		"02/01/06",
		"02/01/2006",
		"02-Jan-06",
	}
	for _, lo := range layouts {
		if tm, err := time.Parse(lo, s); err == nil {
			return tm, true
		}
	}
	return time.Time{}, false
}

func TryToTime(s string) (time.Time, bool) {
	var layouts = []string{
		"15:04:05",
		"3:04:05PM",
		"3:04:05 PM",
		"3:04:05pm",
		"3:04:05 pm",
		"3:04:05 P.M.",
		"3:04:05 p.m.",
	}
	for _, lo := range layouts {
		if tm, err := time.Parse(lo, s); err == nil {
			return tm, true
		}
	}
	return time.Time{}, false
}
