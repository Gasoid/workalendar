package greece

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"3/25":  core.Event("Independence Day"),
		"10/28": core.Event("Ohi Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.EnableOrthodox(),
		core.WithLabourDay(),
		core.WithNewYearDay(),
		core.WithEpiphany(),
		core.WithCleanMonday(),
		core.WithGoodFriday(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithWhitSunday(),
		core.WithWhitMonday(),
		core.WithAssumption(),
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
