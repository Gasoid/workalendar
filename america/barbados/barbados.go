// Code generated github.com/Gasoid/workalendar/generator/gen.go DO NOT EDIT.
package barbados

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
       "1/21": core.Event("Errol Barrow Day"),
       "4/28": core.Event("National Heroes Day"),
       "11/30": core.Event("Independance Day"),
       "12/26": core.Event("Boxing Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithGoodFriday(),
		core.WithEasterMonday(),
		core.WithEasterSunday(),
		core.WithWhitMonday(),
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
		return nil, fmt.Errorf("there is no holiday for %s", date)
	}
	holiday := calendar.GetHoliday(date)
	if holiday == nil {
		return nil, fmt.Errorf("there is no holiday for %s", date)
	}
	return calendar.GetHoliday(date), nil
}
