package slovenia

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"2/8":   core.Event("Preseren Day, the Slovenian Cultural Holiday"),
		"4/27":  core.Event("Day of Uprising Against Occupation"),
		"5/2":   core.Event("Labour Day"),
		"6/25":  core.Event("Statehood Day"),
		"10/31": core.Event("Reformation Day"),
		"11/1":  core.Event("Day of Remembrance of the Dead"),
		"12/26": core.Event("Independence and Unity Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithWhitSunday(),
		core.WithAssumption(),
		core.WithChristmas(),
		core.WithHolidayFunc(variableHolidays),
	)
)

func variableHolidays(date time.Time) core.Holidays {
	days := core.Holidays{}
	year := date.Year()
	if (1955 <= year && year <= 2012) || year >= 2017 {
		days[fmt.Sprintf("%d/1/2", year)] = core.Event("January 2nd")
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
