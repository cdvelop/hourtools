package timetools

func MonthDays(year int) map[int]int {

	var feb_days = 28

	if IsLeap(year) {
		feb_days = 29
	}

	return map[int]int{
		1:  31,
		2:  feb_days,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
}

func SpanishMonth() map[int]string {
	return map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}
}

func ShortMonth() map[int]string {
	return map[int]string{
		1:  "Ene",
		2:  "Feb",
		3:  "Mar",
		4:  "Abr",
		5:  "May",
		6:  "Jun",
		7:  "Jul",
		8:  "Ago",
		9:  "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dic",
	}
}
