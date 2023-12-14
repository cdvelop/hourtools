package timetools

import (
	"github.com/cdvelop/strings"
)

// real_date:2006-01-02 15:04:05 fake_date 2006-01-02 or same real date
// return: date:"2006-01-02" hour:"15:04" optional seconds:"15:04:05"
func DateToDayHour(real_date, fake_date string, seconds ...bool) (date, hour string) {

	var sec string

	if fake_date != "" {
		date, hour, sec = SeparateDateHourAndSeconds(fake_date)

		// si fake_date no contiene la hora se la agregamos de la original
		if strings.Contains(fake_date, ":") == 0 {
			_, hour, sec = SeparateDateHourAndSeconds(real_date)
		}

	} else {

		date, hour, sec = SeparateDateHourAndSeconds(real_date)

	}

	var with_sec string

	for _, with := range seconds {
		if with {
			with_sec = ":" + sec
		}
	}

	hour += with_sec

	return
}
