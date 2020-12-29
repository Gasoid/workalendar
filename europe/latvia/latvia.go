package latvia

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: variable holidays
var (
	holidays = core.Holidays{
		"6/23":  core.Event("Midsummer Day"),
		"6/24":  core.Event("St. John's Day"),
		"11/18": core.Event("Proclamation Day"),
		"12/31": core.Event("New Years Eve"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithGoodFriday(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithChristmasEve(),
		core.WithChristmas(),
		core.WithBoxingDay(),
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
