package timetools_test

import (
	"testing"

	"github.com/cdvelop/timetools"
)

func TestContains(t *testing.T) {

	var testCases = map[string]struct {
		text     string
		search   string
		expected int
	}{
		"Caso1": {
			text:     "Hola, mundo!",
			search:   "mundo",
			expected: 1,
		},
		"Caso2": {
			text:     "Hola, mundo!",
			search:   "golang",
			expected: 0,
		},
		"Caso3": {
			text:     "Hola, mundo!",
			search:   "",
			expected: 0,
		},
		"Caso4": {
			text:     "Hola",
			search:   "Hola, mundo!",
			expected: 0,
		},
		"Caso5": {
			text:     "abracadabra",
			search:   "abra",
			expected: 2,
		},
		"Caso6": {
			text:     "abracadabra",
			search:   "bra",
			expected: 2,
		},
		"Caso7": {
			text:     "abra,cadabra",
			search:   ",",
			expected: 1,
		},
		"Caso8": {
			text:     "(abraLol,*?¡¡",
			search:   "Lol",
			expected: 1,
		},
		"Caso9 ": {
			text:     "(abraLol,*?¡¡",
			search:   "LoL",
			expected: 0,
		},
		"Caso10 ": {
			text:     "(¡ab¡raLol,*?¡¡",
			search:   "¡",
			expected: 4,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {

			expected := timetools.Contains(tc.text, tc.search)
			if expected != tc.expected {
				t.Errorf("Error: Se esperaba %v, pero se obtuvo %v. Texto: %s, Búsqueda: %s", tc.expected, expected, tc.text, tc.search)
			}
		})
	}
}
