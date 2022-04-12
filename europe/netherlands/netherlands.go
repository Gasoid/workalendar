package netherlands

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: add school holidays
var (
	holidays = core.Holidays{
		"5/5": core.Event("Liberation Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithGoodFriday(),
		core.WithEasterMonday(),
		core.WithEasterSunday(),
		core.WithAscension(),
		core.WithWhitSunday(),
		core.WithWhitMonday(),
		core.WithBoxingDay(),
	)
)

func variableHolidays(date time.Time) core.Holidays {
	year := date.Year()
	holidays := core.Holidays{}
	if year > 2013 {
		kingDay := time.Date(year, 4, 27, 0, 0, 0, 0, time.UTC)
		if kingDay.Weekday() != time.Sunday {
			holidays[kingDay.Format("2006/1/2")] = core.Event("King's day")
		} else {
			kingDay = kingDay.AddDate(0, 0, -1)
			holidays[kingDay.Format("2006/1/2")] = core.Event("King's day")
		}
	} else {
		queenDay := time.Date(year, 4, 30, 0, 0, 0, 0, time.UTC)
		if queenDay.Weekday() != time.Sunday {
			holidays[queenDay.Format("2006/1/2")] = core.Event("Queen's day")
		} else {
			queenDay = queenDay.AddDate(0, 0, -1)
			holidays[queenDay.Format("2006/1/2")] = core.Event("Queen's day")
		}
	}
	// start := core.EasterSunday(year).AddDate(0, 0, -7*7)
	// for i := 0; i < 3; i++ {
	// 	day := start.AddDate(0, 0, i)
	// 	holidays[day.Format("2006/1/2")] = core.Event("Carnival")
	// }
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

//NextHoliday returns next holiday from provided date
func NextHoliday(date time.Time) *core.CalEvent {
	return calendar.NextHoliday(date)
}
