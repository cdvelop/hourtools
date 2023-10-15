package timetools

// verifica y cambia formato de fecha según separador ingresado
// ej:2006-01-02 "/" = 2006/01/02
func ChangeDateSeparator(date, separator string) (new_date string, err error) {

	err = CorrectFormatDate(date)
	if err != nil {
		return "", err
	}
	//         year                   month                   day
	return date[:4] + separator + date[5:7] + separator + date[8:10], nil
}
