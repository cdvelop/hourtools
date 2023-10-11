package timetools

func SpanishMonth() *map[int]string {
	s := map[int]string{
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
	return &s
}

func ShortMonth() map[int]string {
	s := map[int]string{
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
	return s
}
