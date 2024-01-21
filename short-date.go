package timetools

import (
	"strconv"
)

type TimeWeek interface {
	//date ej: 2006-01-02,  0 para Domingo, 1 para Lunes, etc.
	WeekDayNumber(date_in string) (n int, err string)
}

// formato fecha "2006-01-02" out; Vie,25,Dic,2011
func StringDateToShort(date_in string, week TimeWeek) (textDay, day, textMonth, year, err string) {
	const this = "StringDateToShort "
	err = CorrectFormatDate(date_in)
	if err != "" {
		err = this + err
		return
	}

	y, m, d, err := stringToDateNumberSeparate(date_in)
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
