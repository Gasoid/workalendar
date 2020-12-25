package belgium

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"7/21":  core.Event("National Day"),
		"11/11": core.Event("Armistice of 1918"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithChristmas(),
		core.WithEasterMonday(),
		core.WithWhitMonday(),
		core.WithAscension(),
		core.WithAllSaints(),
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
		return nil, fmt.Errorf("There is no holiday for %s", date)
	}
	holiday := calendar.GetHoliday(date)
	if holiday == nil {
		return nil, fmt.Errorf("There is no holiday for %s", date)
	}
	return calendar.GetHoliday(date), nil
}
