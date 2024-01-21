package timetools_test

import (
	"strconv"
	"testing"

	"github.com/cdvelop/timetools"
)

type timeNow interface {
	//ej true: "30-12-2001", default false: "2006-01-02"
	DateToDay(DateFormatStructPointer any) string
	//ej default: "2006-01-02", "15:04" WithSeconds true = 15:04:05
	DateToDayHour(DateFormatStructPointer any) (date, hour string)
}

type today struct {
	hour
}

func (today) DateToDay(DateFormatStructPointer any) string {
	return "2023-06-21"
}

type yesterday struct {
	hour
}

func (yesterday) DateToDay(DateFormatStructPointer any) string {
	return "2023-02-20"
}

type tomorrowIsLeap struct {
	hour
}

func (tomorrowIsLeap) DateToDay(DateFormatStructPointer any) string {
	return "2024-02-29"
}

type hour struct{}

func (hour) DateToDayHour(DateFormatStructPointer any) (date, hour string) {
	return "2006-01-02", "15:04:05"
}

func Test_HowOldAreYou(t *testing.T) {

	var (
		dataDateShort = map[string]struct {
			birthDay    string
			timeHandler timeNow
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

			if err != "" {
				ageResp = err
			}

			// and error handler
			if ageResp != data.ageExpected {
				t.Fatalf("*error\n-se esperaba: [%v]\n-pero se obtuvo: [%v]\n-en la entrada: [%v]", data.ageExpected, ageResp, data.birthDay)
			}

		})
	}
}
