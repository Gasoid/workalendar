package iceland

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: variable holidays

var (
	holidays = core.Holidays{
		"6/17":  core.Event("Icelandic National Day"),
		"12/31": core.Event("New Year's Eve"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithLabourDay(),
		core.WithNewYearDay(),
		core.WithHolyThursday(),
		core.WithGoodFriday(),
		core.WithEasterMonday(),
		//core.WithEasterSunday(),
		core.WithAscension(),
		core.WithWhitMonday(),
		core.WithChristmasEve(),
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

//NextHoliday returns next holiday from provided date
func NextHoliday(date time.Time) *core.CalEvent {
	return calendar.NextHoliday(date)
}
