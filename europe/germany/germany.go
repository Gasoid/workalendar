package germany

import (
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{}
	calendar = core.Calendar{
		Days: holidays,
	}
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
func GetHoliday(date time.Time) *core.CalEvent {
	return calendar.GetHoliday(date)
}
