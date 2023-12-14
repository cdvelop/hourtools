package timetools

import (
	"strconv"
)

// date_in: "2006-01-02 15:04:05" or "2006-01-02"
// out: date:"2006-01-02" hour:"15:04" seconds:"05"
func SeparateDateHourAndSeconds(date_in string) (date, hour, seconds string) {

	if len(date_in) == 19 {
		// Formato "2006-01-02 15:04:05"
		date = date_in[:10]
		hour = date_in[11:16]
		seconds = date_in[17:]
	} else if len(date_in) == 10 {
		// Formato "2006-01-02"
		date = date_in
		hour = "00:00"
		seconds = "00"
	}

	return
}

// formato fecha "2006-01-02" retorna: 2006,1,2. NOTA: NO VERIFICA EL FORMATO INGRESADO
func stringToDateSeparate(date string) (year, month, day int, err string) {

	//YEAR
	year, e := strconv.Atoi(date[:4])
	if e != nil {
		err = e.Error()
		return
	}

	//MONTH
	monthText := date[5:7]

	if monthText >= "01" && monthText <= "09" {
		month, e = strconv.Atoi(string(monthText[1]))
	} else if monthText >= "10" && monthText <= "12" {
		month, e = strconv.Atoi(monthText)
	} else {
		err = "error " + monthText + " es un formato de mes incorrecto ej: 01 a 12"
		return
	}
	if e != nil {
		err = e.Error()
		return
	}

	//DAY
	dayTxt := date[8:10]

	if dayTxt >= "01" && dayTxt <= "09" {
		day, e = strconv.Atoi(string(dayTxt[1]))
	} else if dayTxt >= "10" && dayTxt <= "31" {
		day, e = strconv.Atoi(dayTxt)
	}

	if e != nil {
		err = e.Error()
		return
	}

	return
}
