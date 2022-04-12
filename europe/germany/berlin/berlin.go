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
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithHolidayFunc(additionalHolidays),
		core.WithChristmas(),
		core.WithBoxingDay(),
		core.WithEasterMonday(),
		core.WithWhitMonday(),
		core.WithAscension(),
		core.WithGoodFriday(),
	)
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
