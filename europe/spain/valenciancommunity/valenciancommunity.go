package valenciancommunity

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"10/12": core.Event("Fiesta nacional de España"),
		"12/6":  core.Event("Día de la Constitución Española"),
		"3/19":  core.Event("San José"),
		"10/9":  core.Event("Dia de la Comunitat Valenciana"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithEpiphany(),
		core.WithGoodFriday(),
		core.WithAssumption(),
		core.WithAllSaints(),
		core.WithImmaculateConception(),
		core.WithEasterMonday(),
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
