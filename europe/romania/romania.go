package romania

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"1/2":   core.Event("Day After New Year"),
		"1/24":  core.Event("Union Day"),
		"8/15":  core.Event("Dormition of the Theotokos"),
		"11/30": core.Event("St. Andrew's Day"),
		"12/1":  core.Event("National Day/Great Union"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.EnableOrthodox(),
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithGoodFriday(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithWhitSunday(),
		core.WithWhitMonday(),
		core.WithChristmas(),
		core.WithBoxingDay(),
		core.WithHolidayFunc(variableHolidays),
	)
)

func variableHolidays(date time.Time) core.Holidays {
	days := core.Holidays{}
	year := date.Year()
	if year >= 2017 {
		days[fmt.Sprintf("%d/6/1", year)] = core.Event("Children's Day")
	}
	if 1949 <= year && year <= 1990 {
		days[fmt.Sprintf("%d/8/23", year)] = core.Event("Liberation from Fascist Occupation Day")
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
