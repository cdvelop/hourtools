package timetools

import (
	"strconv"

	"github.com/cdvelop/model"
)

// verifica formato y valores numéricos en sus posiciones ej: "2006-01-02"
func CorrectFormatDate(date string) error {

	if len(date) != 10 {
		return model.Error("formato de fecha ingresado incorrecto ej: 2006-01-02")
	}

	numMap := map[byte]bool{
		'0': true, '1': true, '2': true, '3': true, '4': true,
		'5': true, '6': true, '7': true, '8': true, '9': true,
	}

	// Verificar que los guiones estén en las posiciones correctas y que los caracteres sean números
	for i, char := range date {
		if i == 4 || i == 7 {
			if char != '-' {
				return model.Error("formato de fecha ingresado incorrecto ej: 2006-01-02")
			}
		} else {
			if !numMap[byte(char)] {
				return model.Error("formato de fecha ingresado incorrecto ej: 2006-01-02")
			}
		}
	}

	return nil
}

// verifica formato 2006-01-02 y si los rangos de el año, mes y dia son validos
// y si los Dias existen según año y mes bisiesto
func CheckDateExists(date string) error {

	err := CorrectFormatDate(date)
	if err != nil {
		return err
	}

	year, month, day := stringToDateSeparate(date)

	// Verificar los rangos para año, mes y día

	if year < 1000 || year > 9999 {
		return model.Error("año fuera de rango")
	}

	if month < 1 || month > 12 {
		return model.Error("mes fuera de rango")
	}

	month_days := MonthDays(year)[month]
	if day < 1 {
		return model.Error("día no puede ser cero")
	}

	if day > month_days {

		msg := SpanishMonth()[month] + " no contiene " + strconv.Itoa(day) + " días."

		if !IsLeap(year) && month == 2 {
			msg += " año " + date[:4] + " no es bisiesto."
		}

		return model.Error(msg)
	}

	return nil
}
