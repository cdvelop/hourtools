package timetools

type TimeNow interface {
	//ej: 2006-01-02
	DateToDay() string
	//ej: 2006-01-02 15:04:05
	DateToDayHour() string
}

type TimeWeek interface {
	//date ej: 2006-01-02,  0 para Domingo, 1 para Lunes, etc.
	WeekDayNumber(date_in string) (n int, err string)
}
