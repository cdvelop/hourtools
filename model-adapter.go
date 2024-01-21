package timetools

type TimeNow interface {
	//ej true: "30-12-2001", default false: "2006-01-02"
	DateToDay(DateFormatAdapter) string
	//ej default: "2006-01-02", "15:04" WithSeconds true = 15:04:05
	DateToDayHour(DateFormatAdapter) (date, hour string)
}
