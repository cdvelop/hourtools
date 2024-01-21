package timetools

import (
	"strconv"
)

// verifica formato y valores numéricos en sus posiciones ej: "2006-01-02"
func CorrectFormatDate(date string) (err string) {

	if len(date) != 10 {
		return "formato de fecha ingresado incorrecto ej: 2006-01-02"
	}

	numMap := map[byte]bool{
		'0': true, '1': true, '2': true, '3': true, '4': true,
		'5': true, '6': true, '7': true, '8': true, '9': true,
	}

	// Verificar que los guiones estén en las posiciones correctas y que los caracteres sean números
	for i, char := range date {
		if i == 4 || i == 7 {
			if char != '-' {
				return "formato de fecha ingresado incorrecto ej: 2006-01-02"
			}
		} else {
			if !numMap[byte(char)] {
				return "formato de fecha ingresado incorrecto ej: 2006-01-02"
			}
		}
	}

	return ""
}

// verifica formato 2006-01-02 y si los rangos de el año, mes y dia son validos
// y si los Dias existen según año y mes bisiesto
func CheckDateExists(date string) (err string) {
	const this = "CheckDateExists "
	err = CorrectFormatDate(date)
	if err != "" {
		return this + err
	}

	year, month, day, er := stringToDateNumberSeparate(date)
	if er != "" {
		err = this + er
		return
	}

	// Verificar los rangos para año, mes y día

	if year < 1000 || year > 9999 {
		return "año fuera de rango"
	}

	if month < 1 || month > 12 {
		return "mes fuera de rango"
	}

	month_days := MonthDays(year)[month]
	if day < 1 {
		return "día no puede ser cero"
	}

	if day > month_days {

		err = SpanishMonth()[month] + " no contiene " + strconv.Itoa(day) + " días."

		if !IsLeap(year) && month == 2 {
			err += " año " + date[:4] + " no es bisiesto."
		}

		return err
	}

	return ""
}
