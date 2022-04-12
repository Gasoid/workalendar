package france

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: check christmas
var (
	holidays = core.Holidays{
		"5/8":   core.Event("Victory in Europe Day"),
		"7/14":  core.Event("Bastille Day"),
		"11/11": core.Event("Armistice Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithEasterMonday(),
		core.WithAscension(),
		core.WithWhitMonday(),
		core.WithAllSaints(),
		core.WithAssumption(),
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
