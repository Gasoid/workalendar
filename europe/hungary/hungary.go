package hungary

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: add good friday as a variable holiday
var (
	holidays = core.Holidays{
		"3/15":  core.Event("National Day"),
		"8/20":  core.Event("St Stephen's Day"),
		"10/23": core.Event("National Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithLabourDay(),
		core.WithNewYearDay(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithWhitSunday(),
		core.WithWhitMonday(),
		core.WithBoxingDay(),
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
