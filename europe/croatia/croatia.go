package croatia

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"1/2":  core.Event("Day After New Year"),
		"3/8":  core.Event("International Women's Day"),
		"5/9":  core.Event("Victory Day"),
		"7/3":  core.Event("Republic Day"),
		"11/7": core.Event("October Revolution Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithEpiphany(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithCorpusChristi(),
		core.WithAssumption(),
		core.WithAllSaints(),
		core.WithChristmas(),
		core.WithBoxingDay(),
		core.WithHolidayFunc(additionalHolidays),
	)
)

func additionalHolidays(date time.Time) core.Holidays {
	if date.Year() >= 2020 {
		days := core.Holidays{
			fmt.Sprintf("%d/5/30", date.Year()):  core.Event("Statehood Day"),
			fmt.Sprintf("%d/11/18", date.Year()): core.Event("Remembrance Day"),
		}
		return days
	}
	days := core.Holidays{
		fmt.Sprintf("%d/6/25", date.Year()): core.Event("Statehood Day"),
		fmt.Sprintf("%d/10/8", date.Year()): core.Event("Independence Day"),
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
	return holiday, nil
}
