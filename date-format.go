package timetools

type DateFormatAdapter interface {
	LeftDay() bool     //true ej: "30-12-2001", default false: "2006-01-02"
	WithSeconds() bool //ej: 2006-01-02, 15:04, WithSeconds true = 15:04:05
}

type DateFormat struct {
	Left_Day     bool //true ej: "30-12-2001", default false: "2006-01-02"
	With_Seconds bool //ej: 2006-01-02, 15:04, WithSeconds true = 15:04:05
}

func (d DateFormat) LeftDay() bool {
	return d.Left_Day
}

func (d DateFormat) WithSeconds() bool {
	return d.With_Seconds
}
