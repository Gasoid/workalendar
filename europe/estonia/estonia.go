package estonia

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"2/24": core.Event("Independence Day"),
		"5/1":  core.Event("Kevadpüha"),
		"6/23": core.Event("Võidupüha"),
		"6/24": core.Event("Jaanipäev"),
		"8/20": core.Event("Taasiseseisvumispäev"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithChristmas(),
		core.WithEasterSunday(),
		core.WithWhitSunday(),
		core.WithChristmasEve(),
		core.WithGoodFriday(),
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
