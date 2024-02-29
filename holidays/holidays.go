package holidays

import (
	"math"
	"reflect"
	"time"
)

type HolidaysDate struct {
	Pascoa               time.Time
	Carnaval             time.Time
	SextaFeiraSanta      time.Time
	CorpusChristi        time.Time
	AnoNovo              time.Time
	Tiradentes           time.Time
	DiaDoTrabalhador     time.Time
	Independencia        time.Time
	NossaSraAparecida    time.Time
	Finados              time.Time
	ProclamacaoRepublica time.Time
	Natal                time.Time
}

func Holidays(intialDate time.Time) HolidaysDate {
	easter := holidayEaster(intialDate)

	return HolidaysDate{
		Pascoa:               easter,
		Carnaval:             holidayCarnival(easter),
		SextaFeiraSanta:      holidayGodFriday(easter),
		CorpusChristi:        calcCorpusChristi(easter),
		AnoNovo:              time.Date(intialDate.Year(), time.January, 1, 0, 0, 0, 0, time.UTC),
		Tiradentes:           time.Date(intialDate.Year(), time.April, 21, 0, 0, 0, 0, time.UTC),
		DiaDoTrabalhador:     time.Date(intialDate.Year(), time.May, 1, 0, 0, 0, 0, time.UTC),
		Independencia:        time.Date(intialDate.Year(), time.September, 7, 0, 0, 0, 0, time.UTC),
		NossaSraAparecida:    time.Date(intialDate.Year(), time.October, 12, 0, 0, 0, 0, time.UTC),
		Finados:              time.Date(intialDate.Year(), time.November, 2, 0, 0, 0, 0, time.UTC),
		ProclamacaoRepublica: time.Date(intialDate.Year(), time.November, 15, 0, 0, 0, 0, time.UTC),
		Natal:                time.Date(intialDate.Year(), time.December, 25, 0, 0, 0, 0, time.UTC),
	}
}

func IsHoliday(intialDate time.Time) bool {
	compareDate := time.Date(intialDate.Year(), intialDate.Month(), intialDate.Day(), 0, 0, 0, 0, time.UTC)
	holidays := Holidays(compareDate)

	holidaysValue := reflect.ValueOf(holidays)
	for i := 0; i < holidaysValue.NumField(); i++ {
		holidayDate := holidaysValue.Field(i).Interface().(time.Time)

		if compareDate.Equal(holidayDate) {
			return true
		}
	}

	return false
}

// Method to calculate when easter will occurs.
// Gauss Algorithm used in this function "https://pt.wikipedia.org/wiki/C%C3%A1lculo_da_P%C3%A1scoa"
func holidayEaster(actualDate time.Time) time.Time {
	var x float64
	var y float64
	dateEaster := time.Now().UTC()
	year := float64(actualDate.Year())

	switch {
	case year >= 1582 && year <= 1699:
		x = 22
		y = 2
	case year >= 1700 && year <= 1799:
		x = 23
		y = 3
	case year >= 1800 && year <= 1899:
		x = 23
		y = 4
	case year >= 1900 && year <= 2099:
		x = 24
		y = 5
	case year >= 2100 && year <= 2199:
		x = 24
		y = 6
	case year >= 2200 && year <= 2299:
		x = 25
		y = 7
	}

	a := math.Mod(year, 19)
	b := math.Mod(year, 4)
	c := math.Mod(year, 7)
	d := math.Mod((19*a)+x, 30)
	e := math.Mod((2*b)+(4*c)+(6*d)+y, 7)

	if (d + e) < 10 {
		dateEaster = time.Date(actualDate.Year(), 3, (int)(d+e+22), 0, 0, 0, 0, time.UTC)
	} else {
		dateEaster = time.Date(actualDate.Year(), 4, (int)(d+e-9), 0, 0, 0, 0, time.UTC)

		if dateEaster.Day() == 26 || (dateEaster.Day() == 25 && d == 28 && a > 10) {
			dateEaster = time.Date(actualDate.Year(), 4, (int)(d+e-9-7), 0, 0, 0, 0, time.UTC)
		}
	}

	return dateEaster
}

// Method to calculate carnival based on easter date "https://pt.wikipedia.org/wiki/Carnaval#C%C3%A1lculo"
func holidayCarnival(date time.Time) time.Time {
	return date.AddDate(0, 0, -47)
}

// Method to calculate God Friday based on easter date "https://pt.wikipedia.org/wiki/Sexta-feira_Santa#Data"
func holidayGodFriday(date time.Time) time.Time {
	return date.AddDate(0, 0, -2)
}

// Method to calculate Corpus Christi based on easter date
func calcCorpusChristi(date time.Time) time.Time {
	return date.AddDate(0, 0, 60)
}
