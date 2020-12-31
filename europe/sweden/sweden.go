package sweden

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: add variable holidays
var (
	holidays = core.Holidays{
		"6/6":   core.Event("National Day"),
		"12/31": core.Event("New Year's Eve"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithEpiphany(),
		core.WithGoodFriday(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithAscension(),
		core.WithWhitSunday(),
		core.WithChristmasEve(),
		core.WithBoxingDay(),
	)
)

func variableHolidays(date time.Time) core.Holidays {
	days := core.Holidays{}
	return days
}

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
