package berlin

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"10/3": core.Event("Day of German Unity"),
	}
	calendar = core.Calendar(holidays, additionalHolidays)
)

func additionalHolidays(date time.Time) core.Holidays {
	days := make(core.Holidays, 1)
	if date.Year() >= 2019 {
		days[fmt.Sprintf("%d/3/8", date.Year())] = core.Event("International Women's Day")
	}
	if date.Year() == 2020 {
		days[fmt.Sprintf("%d/5/8", date.Year())] = core.Event("Liberation Day")
	}
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
	// if isShiftHoliday(date) {
	// 	holiday := getOriginalHoliday(date)
	// 	if holiday == nil {
	// 		return nil, fmt.Errorf("There is no holiday for %s", date)
	// 	}
	// 	return holiday, nil
	// }
	holiday := calendar.GetHoliday(date)
	if holiday == nil {
		return nil, fmt.Errorf("There is no holiday for %s", date)
	}
	return calendar.GetHoliday(date), nil
}
