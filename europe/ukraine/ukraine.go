package ukraine

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"3/8": core.Event("International Women’s Day"),
		"5/9": core.Event("Victory Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.EnableOrthodox(),
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithGoodFriday(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithWhitMonday(),
		core.WithHolidayFunc(variableHolidays),
	)
)

func variableHolidays(date time.Time) core.Holidays {
	days := core.Holidays{}
	year := date.Year()
	orthodoxChristmas := time.Date(year, 1, 7, 0, 0, 0, 0, time.UTC)
	orthodoxChristmas = core.FindWorkingDay(orthodoxChristmas)
	days[orthodoxChristmas.Format("2006/1/2")] = core.Event("Orthodox Christmas")

	if year >= 1996 {
		constitutionDay := time.Date(year, 6, 28, 0, 0, 0, 0, time.UTC)
		constitutionDay = core.FindWorkingDay(constitutionDay)
		days[constitutionDay.Format("2006/1/2")] = core.Event("Constitution Day")
	}

	if year == 1991 {
		days["1991/6/16"] = core.Event("Independence Day")
	}

	if year >= 1992 {
		independenceDay := time.Date(year, 8, 24, 0, 0, 0, 0, time.UTC)
		independenceDay = core.FindWorkingDay(independenceDay)
		days[independenceDay.Format("2006/1/2")] = core.Event("Independence Day")
	}
	if year >= 2015 {
		days[fmt.Sprintf("%d/10/14", year)] = core.Event("Day of Defender of Ukraine")
	}

	if year >= 2017 {
		days[fmt.Sprintf("%d/12/25", year)] = core.Event("Christmas Day")
	}

	if year <= 2017 {
		days[fmt.Sprintf("%d/5/2", year)] = core.Event("Workers Solidarity Day")
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
