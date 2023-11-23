package timetools_test

import (
	"testing"

	"github.com/cdvelop/timetools"
)

func TestDayOfYear(t *testing.T) {
	tests := []struct {
		year, month, day  int
		expectedDayOfYear int
		expectedError     string
	}{
		{2023, 3, 31, 90, ""}, // Fecha válida en año no bisiesto
		{2023, 3, 1, 60, ""},  // Fecha válida en año no bisiesto
		{2024, 2, 29, 60, ""}, // Fecha válida en año bisiesto
		{2023, 2, 29, -1, "DayOfYear error Fecha no válida"}, // Fecha inválida en año no bisiesto
		{2024, 4, 31, -1, "DayOfYear error Fecha no válida"}, // Fecha inválida en año bisiesto
	}

	for _, test := range tests {
		dayOfYear, err := timetools.DayOfYear(test.year, test.month, test.day)
		if dayOfYear != test.expectedDayOfYear || (err != test.expectedError) {
			t.Errorf("Para %d/%d/%d, se esperaba %d con error [%v], pero se obtuvo %d con error [%v]", test.year, test.month, test.day, test.expectedDayOfYear, test.expectedError, dayOfYear, err)
		}
	}
}
