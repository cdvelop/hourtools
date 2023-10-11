package timetools

import (
	"github.com/cdvelop/model"
)

// date format ej: "2006-01-02"
func CorrectDate(date string) error {

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
