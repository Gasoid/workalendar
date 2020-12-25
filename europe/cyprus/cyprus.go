package croatia

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"3/25":  core.Event("Greek Independence Day"),
		"4/1":   core.Event("Cyprus National Day"),
		"7/15":  core.Event("Dormition of the Theotokos"),
		"10/1":  core.Event("Cyprus Independence Day"),
		"10/28": core.Event("Greek National Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithEpiphany(),
		core.WithCleanMonday(),
		core.WithGoodFriday(),
		core.WithEasterSaturday(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithEasterTuesday(),
		core.WithWhitMonday(),
		core.WithChristmasEve(),
		core.WithChristmas(),
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
