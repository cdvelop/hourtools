package timetools

import (
	"strconv"

	"github.com/cdvelop/model"
)

// formato fecha "2006-01-02" out; Vie,25,Dic,2011
func StringDateToShort(date_in string, week model.TimeWeek) (textDay, day, textMonth, year string, err error) {

	err = CorrectDate(date_in)
	if err != nil {
		return
	}

	y, m, d, err1 := StringToDateSeparate(date_in)
	if err1 != nil {
		err = err1
		return
	}

	w, err2 := week.WeekDayNumber(date_in)
	if err2 != nil {
		err = err2
		return
	}

	textDay = ShortDay()[w]

	day = strconv.Itoa(d)

	textMonth = ShortMonth()[m]

	year = strconv.Itoa(y % 100)

	return
}
