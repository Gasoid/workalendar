package graubunden

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"8/1": core.Event("National Holiday"),
		//"1/2": core.Event("Berchtold's Day"),
		"3/19": core.Event("St Joseph's Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithGoodFriday(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithAscension(),
		core.WithWhitSunday(),
		core.WithWhitMonday(),
		core.WithChristmas(),
		core.WithBoxingDay(),
		core.WithEpiphany(),
		core.WithCorpusChristi(),
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
