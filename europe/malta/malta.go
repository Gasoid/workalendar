package malta

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: recheck last 3 holidays
var (
	holidays = core.Holidays{
		"3/31":  core.Event("Freedom Day"),
		"6/7":   core.Event("Sette Giugno"),
		"9/8":   core.Event("Victory Day"),
		"9/21":  core.Event("Independence Day"),
		"12/13": core.Event("Republic Day"),
		"2/10":  core.Event("Feast of Saint Paul's Shipwreck"),
		"3/19":  core.Event("Feast of Saint Joseph"),
		"6/29":  core.Event("Feast of Saint Peter & Saint Paul"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithGoodFriday(),
		core.WithAssumption(),
		core.WithImmaculateConception(),
		core.WithChristmasEve(),
		core.WithLabourDay(),
	)
)

func variableHolidays(date time.Time) core.Holidays {
	return nil
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
