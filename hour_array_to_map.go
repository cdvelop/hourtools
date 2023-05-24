package hour

// ej: "08:00" -> "08:00": "8:00"
func HourArrayToMap(hour_in []string) (out map[string]string) {
	out = make(map[string]string, len(hour_in))

	for _, hour_key := range hour_in {

		if len(hour_key) > 1 {
			var value string

			if hour_key[0:1] == "0" {
				value = hour_key[1:]
			} else {
				value = hour_key
			}

			out[hour_key] = value
		}
	}

	return
}
