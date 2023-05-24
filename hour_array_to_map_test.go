package go_hour_test

import (
	"log"
	"testing"

	"github.com/cdvelop/hour"
)

// ej: "08:00" -> "08:00": "8:00"
func Test_HourArrayToMap(t *testing.T) {

	hours := []string{"08:00", "09:00", "13:00"}

	out := hour.HourArrayToMap(hours)

	if len(out) != len(hours) {
		log.Fatal("error salida diferente")
	}
	// fmt.Printf("salida: %v\n", out)
}
