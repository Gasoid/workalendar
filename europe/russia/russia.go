package russia

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"1/2":  core.Event("Day After New Year"),
		"2/23": core.Event("Defendence of the Fatherland"),
		"3/8":  core.Event("International Women's Day"),
		"5/9":  core.Event("Victory Day"),
		"6/12": core.Event("National Day"),
		"11/4": core.Event("Day of Unity"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithHolidayFunc(additionalHolidays),
		core.WithLabourDay(),
	)
)

func additionalHolidays(date time.Time) core.Holidays {
	if date.Year() >= 2005 {
		days := core.Holidays{
			fmt.Sprintf("%d/1/3", date.Year()): core.Event("Third Day after New Year"),
			fmt.Sprintf("%d/1/4", date.Year()): core.Event("Fourth Day after New Year"),
			fmt.Sprintf("%d/1/5", date.Year()): core.Event("Fifth Day after New Year"),
			fmt.Sprintf("%d/1/6", date.Year()): core.Event("Sixth Day after New Year"),
			fmt.Sprintf("%d/1/8", date.Year()): core.Event("Eighth Day after New Year"),
		}
		return days
	}
	return nil
}

func isShiftHoliday(date time.Time) bool {
	if date.Month() == time.January && (date.Day() >= 1 && date.Day() <= 9) {
		return false
	}

	if date.Weekday() == time.Monday {
		return calendar.IsHoliday(date.AddDate(0, 0, -2)) || calendar.IsHoliday(date.AddDate(0, 0, -1))
	}
	return false
}

func getOriginalHoliday(date time.Time) *core.CalEvent {
	sat := date.AddDate(0, 0, -2)
	sun := date.AddDate(0, 0, -1)
	if calendar.IsHoliday(sat) && calendar.GetHoliday(sat) != nil {
		return calendar.GetHoliday(sat)
	}
	if calendar.IsHoliday(sun) && calendar.GetHoliday(sun) != nil {
		return calendar.GetHoliday(sun)
	}
	return nil
}

//IsWorkingDay is inteded to check whether a day is working or not
func IsWorkingDay(date time.Time) bool {
	return calendar.IsWorkingDay(date) && !isShiftHoliday(date)
}

//IsHoliday is inteded to check whether a day is holiday or not
func IsHoliday(date time.Time) bool {
	return calendar.IsHoliday(date) || isShiftHoliday(date)
}

//GetHoliday is inteded to check whether a day is holiday or not
func GetHoliday(date time.Time) (*core.CalEvent, error) {
	if !IsHoliday(date) {
		return nil, fmt.Errorf("There is no holiday for %s", date)
	}
	if isShiftHoliday(date) {
		holiday := getOriginalHoliday(date)
		if holiday == nil {
			return nil, fmt.Errorf("There is no holiday for %s", date)
		}
		return holiday, nil
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
