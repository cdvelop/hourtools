package timetools

import (
	"strconv"

	"github.com/cdvelop/model"
)

// formato fecha "2006-01-02" retorna: 2006,1,2. NOTA: NO VERIFICA EL FORMATO INGRESADO
func stringToDateSeparate(date string) (year, month, day int) {

	//YEAR
	year, err := strconv.Atoi(date[:4])
	if err != nil {
		return
	}

	//MONTH
	monthText := date[5:7]

	if monthText >= "01" && monthText <= "09" {
		month, err = strconv.Atoi(string(monthText[1]))
	} else if monthText >= "10" && monthText <= "12" {
		month, err = strconv.Atoi(monthText)
	} else {
		err = model.Error("error", monthText, "es un formato de mes incorrecto ej: 01 a 12")
	}
	if err != nil {
		return
	}

	//DAY
	dayTxt := date[8:10]

	if dayTxt >= "01" && dayTxt <= "09" {
		day, err = strconv.Atoi(string(dayTxt[1]))
	} else if dayTxt >= "10" && dayTxt <= "31" {
		day, err = strconv.Atoi(dayTxt)
	}

	return
}
