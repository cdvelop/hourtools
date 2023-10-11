package timetools

import (
	"github.com/cdvelop/model"
)

// obtener edad en formato numero ej: (birthday) "2006-01-02"
func HowOldAreYou(birthday string, time model.TimeNow) (int, error) {

	err := CorrectDate(birthday)
	if err != nil {
		return 0, err
	}

	if time == nil {
		return 0, model.Error("error TimeAdapter: 'ToDay(layout string) string' no ingresado")
	}

	// Divide la fecha de nacimiento y la fecha actual en año, mes y día usando la función strconv.Atoi
	yearBirth, monthBirth, dayBirth, err := StringToDateSeparate(birthday)
	if err != nil {
		return 0, err
	}

	// Obtén la fecha actual en el mismo formato que la fecha de nacimiento
	today := time.ToDay("2006-01-02")

	yearNow, monthNow, dayNow, err := StringToDateSeparate(today)
	if err != nil {
		return 0, err
	}

	// Calcula la diferencia en años
	age := yearNow - yearBirth

	dayOfYearBirth, err := DayOfYear(yearBirth, monthBirth, dayBirth)
	if err != nil {
		return 0, err
	}

	// Si la fecha es anterior a la fecha de nacimiento, entonces no han transcurrido tantos años.
	dayOfYearNow, err := DayOfYear(yearNow, monthNow, dayNow)
	if err != nil {
		return 0, err
	}

	birthDay := GetAdjustedBirthDay(yearBirth, dayOfYearBirth, yearNow, dayOfYearNow)

	if dayOfYearNow < birthDay {
		age -= 1
	}

	return age, nil
}

// Obtiene la fecha de nacimiento ajustada para solucionar las diferencias de los años bisiestos.
func GetAdjustedBirthDay(yearBirth, dayOfYearBirth, yearNow, dayOfYearNow int) int {

	if IsLeap(yearBirth) && !IsLeap(yearNow) && dayOfYearBirth >= 60 {
		return dayOfYearBirth - 1
	}

	if IsLeap(yearNow) && !IsLeap(yearBirth) && dayOfYearNow >= 60 {
		return dayOfYearBirth + 1
	}

	return dayOfYearBirth
}
