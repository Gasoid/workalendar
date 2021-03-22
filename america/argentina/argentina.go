package argentina

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"10/26": core.Event("National Holiday"),
		"3/24":  core.Event("Día Nacional de la Memoria por la Verdad y la Justicia"),
		"5/25":  core.Event("Día de la Revolución de Mayo"),
		"6/20":  core.Event("Día Paso a la Inmortalidad del General Manuel Belgrano"),
		"7/9":   core.Event("Día de la Independencia"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithFatTuesday(),
		core.WithGoodFriday(),
		core.WithEasterSaturday(),
		core.WithEasterSunday(),
		core.WithChristmas(),
		core.WithImmaculateConception(),
	)
)

//IsWorkingDay is inteded to check whether a day is working or not
func IsWorkingDay(date time.Time) bool {
	return calendar.IsWorkingDay(date)
}

//IsHoliday is inteded to check whether a day is holiday or not
func IsHoliday(date time.Time) bool {
	return calendar.IsHoliday(date)
}

//GetHoliday is inteded to check whether a day is holiday or not
func GetHoliday(date time.Time) (*core.CalEvent, error) {
	if !IsHoliday(date) {
		return nil, fmt.Errorf("there is no holiday for %s", date)
	}
	holiday := calendar.GetHoliday(date)
	if holiday == nil {
		return nil, fmt.Errorf("there is no holiday for %s", date)
	}
	return calendar.GetHoliday(date), nil
}
