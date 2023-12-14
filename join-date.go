package timetools

// ej: day:01, month:30, year:2003 default result: "2003-30-01"
// with optional spacer ej: / = "2003/30/01"
//Important: input data is not validated
func JointDate(day, month, year string, spacer ...string) string {
	var space = "-"

	for _, v := range spacer {
		space = v
	}

	return year + space + month + space + day
}
