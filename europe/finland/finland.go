package finland

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: add variable holidays: Midsummer's Eve and the rest
var (
	holidays = core.Holidays{
		"12/6": core.Event("Independence Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithEpiphany(),
		core.WithGoodFriday(),
		core.WithNewYearDay(),
		core.WithChristmas(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithAscension(),
		core.WithWhitSunday(),
		core.WithChristmasEve(),
		core.WithBoxingDay(),
		core.WithLabourDay(),
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
