package czechrepublic

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"1/1":   core.Event("Restoration Day of the Independent Czech State"),
		"5/8":   core.Event("Liberation Day"),
		"7/5":   core.Event("Saints Cyril and Methodius Day"),
		"7/6":   core.Event("Jan Hus Day"),
		"9/28":  core.Event("St. Wenceslas Day (Czech Statehood Day)"),
		"10/28": core.Event("Independent Czechoslovak State Day"),
		"11/17": core.Event("Struggle for Freedom and Democracy Day"),
		"12/26": core.Event("St. Stephen's Day (The Second Christmas Day)"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithEasterMonday(),
		//TODO: check whether good friday is ok
		core.WithGoodFriday(),
		core.WithChristmas(),
		core.WithChristmasEve(),
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
