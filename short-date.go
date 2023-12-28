package timetools

import (
	"strconv"

	"github.com/cdvelop/model"
)

// formato fecha "2006-01-02" out; Vie,25,Dic,2011
func StringDateToShort(date_in string, week model.TimeWeek) (textDay, day, textMonth, year, err string) {
	const this = "StringDateToShort "
	err = CorrectFormatDate(date_in)
	if err != "" {
		err = this + err
		return
	}

	y, m, d, err := stringToDateSeparate(date_in)
	if err != "" {
		err = this + err
		return
	}

	w, err := week.WeekDayNumber(date_in)
	if err != "" {
		err = this + err
		return
	}

	textDay = ShortDay()[w]

	day = strconv.Itoa(d)

	textMonth = ShortMonth()[m]

	year = strconv.Itoa(y % 100)

	return
}
