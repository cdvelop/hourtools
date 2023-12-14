package timetools

// obtener edad en formato numero ej: (birthday) "2006-01-02"
func HowOldAreYou(birthday string, time TimeNow) (age int, err string) {
	const this = "HowOldAreYou "
	err = CorrectFormatDate(birthday)
	if err != "" {
		return 0, this + err
	}

	if time == nil {
		return 0, this + "TimeAdapter: 'DateToDay(layout string) string' no ingresado"
	}

	// Divide la fecha de nacimiento y la fecha actual en año, mes y día usando la función strconv.Atoi
	yearBirth, monthBirth, dayBirth, er := stringToDateSeparate(birthday)
	if er != "" {
		err = this + er
		return
	}

	// Obtén la fecha actual en el mismo formato que la fecha de nacimiento
	today := time.DateToDay()

	yearNow, monthNow, dayNow, er := stringToDateSeparate(today)
	if er != "" {
		err = this + er
		return
	}
	// Calcula la diferencia en años
	age = yearNow - yearBirth

	dayOfYearBirth, err := DayOfYear(yearBirth, monthBirth, dayBirth)
	if err != "" {
		return 0, this + err
	}

	// Si la fecha es anterior a la fecha de nacimiento, entonces no han transcurrido tantos años.
	dayOfYearNow, err := DayOfYear(yearNow, monthNow, dayNow)
	if err != "" {
		return 0, this + err
	}

	birthDay := GetAdjustedBirthDay(yearBirth, dayOfYearBirth, yearNow, dayOfYearNow)

	if dayOfYearNow < birthDay {
		age -= 1
	}

	return age, ""
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
