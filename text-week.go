package timetools

func SpanishDay() *map[int]string {
	s := map[int]string{
		1: "Lunes",
		2: "Martes",
		3: "Miercoles",
		4: "Jueves",
		5: "Viernes",
		6: "Sabado",
		0: "Domingo",
	}
	return &s
}

func ShortDay() map[int]string {
	s := map[int]string{
		1: "Lun",
		2: "Mar",
		3: "Mie",
		4: "Jue",
		5: "Vie",
		6: "Sab",
		0: "Dom",
	}
	return s
}
