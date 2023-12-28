package timetools

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

// real_date:2006-01-02 15:04:05 fake_date 2006-01-02 or same real date
// return: date:"2006-01-02" hour:"15:04" optional seconds:"15:04:05"
// left
func DateToDayHour(real_date, fake_date string, df *model.DateFormat) (date, hour string) {

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
	if df != nil {

		if df.LeftDay {
			changeLeftDayFormat(&date)
		}

		if df.WithSeconds {
			with_sec = ":" + sec
		}
	}

	hour += with_sec

	return
}

// date_in: 2006-01-02 change to: 02-01-2006
// separation ej: "/" default "-"
func changeLeftDayFormat(date *string, separation ...string) {
	if date != nil && len(*date) == 10 {

		dayTxt := (*date)[8:10]

		monthText := (*date)[5:7]

		yearText := (*date)[:4]

		var sep = "-"
		for _, v := range separation {
			sep = v
		}

		*date = dayTxt + sep + monthText + sep + yearText

	}
}
