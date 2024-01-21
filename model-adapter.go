package timetools

type TimeNow interface {
	//ej true: "30-12-2001", default false: "2006-01-02"
	DateToDay(DateFormatStructPointer any) string
}

type DateFormat struct {
	LeftDay     bool //true ej: "30-12-2001", default false: "2006-01-02"
	WithSeconds bool //ej: 2006-01-02, 15:04, WithSeconds true = 15:04:05
}
