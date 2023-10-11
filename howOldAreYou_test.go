package timetools_test

import (
	"strconv"
	"testing"

	"github.com/cdvelop/model"
	"github.com/cdvelop/timetools"
)

type today struct{}

func (today) ToDay(layout string) string {
	return "2023-06-21"
}

type yesterday struct{}

func (yesterday) ToDay(layout string) string {
	return "2023-02-20"
}

type tomorrowIsLeap struct{}

func (tomorrowIsLeap) ToDay(layout string) string {
	return "2024-02-29"
}

func Test_HowOldAreYou(t *testing.T) {

	var (
		dataDateShort = map[string]struct {
			birthDay    string
			timeHandler model.TimeNow
			ageExpected string //31
		}{
			"fecha actual 2023-06-21":              {"1981-06-21", today{}, "42"},
			"fecha actual 2023-02-20":              {"1981-06-21", yesterday{}, "41"},
			"año bisiesto fecha actual 2024-02-29": {"1991-02-28", tomorrowIsLeap{}, "33"},
		}
	)

	for prueba, data := range dataDateShort {
		t.Run((prueba), func(t *testing.T) {

			resp, err := timetools.HowOldAreYou(data.birthDay, data.timeHandler)

			ageResp := strconv.Itoa(resp)

			if err != nil {
				ageResp = err.Error()
			}

			// and error handler
			if ageResp != data.ageExpected {
				t.Fatalf("*error\n-se esperaba: [%v]\n-pero se obtuvo: [%v]\n-en la entrada: [%v]", data.ageExpected, ageResp, data.birthDay)
			}

		})
	}
}
