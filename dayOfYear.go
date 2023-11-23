package timetools

// DayOfYear calcula el día del año para una fecha dada y retorna un error si la fecha no es válida.
func DayOfYear(year, month, day int) (d int, err string) {
	// Definir días en cada mes
	daysInMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Añadir un día al mes de febrero si el año es bisiesto
	if IsLeap(year) {
		daysInMonth[2] = 29
	}

	// Validar mes y día
	if month < 1 || month > 12 || day < 1 || day > daysInMonth[month] {
		return -1, "DayOfYear error Fecha no válida"
	}

	// Calcular el día del año
	dayOfYear := day
	for i := 1; i < month; i++ {
		dayOfYear += daysInMonth[i]
	}

	return dayOfYear, ""
}
