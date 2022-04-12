package slovakia

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"1/1":   core.Event("Day of the Establishment of the Slovak Republic"),
		"5/8":   core.Event("Liberation Day"),
		"7/5":   core.Event("Saints Cyril and Methodius Day"),
		"8/29":  core.Event("Slovak National Uprising anniversary"),
		"9/1":   core.Event("Day of the Constitution of the Slovak Republic"),
		"9/15":  core.Event("Day of Blessed Virgin Mary, patron saint of Slovakia"),
		"11/17": core.Event("Struggle for Freedom and Democracy Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithGoodFriday(),
		core.WithEasterMonday(),
		core.WithEpiphany(),
		core.WithAllSaints(),
		core.WithChristmasEve(),
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

//NextHoliday returns next holiday from provided date
func NextHoliday(date time.Time) *core.CalEvent {
	return calendar.NextHoliday(date)
}
