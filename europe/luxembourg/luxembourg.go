package luxembourg

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"6/23": core.Event("Luxembourg National Holiday"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithEasterMonday(),
		core.WithAscension(),
		core.WithWhitMonday(),
		core.WithAllSaints(),
		core.WithAssumption(),
		core.WithBoxingDay(),
		core.WithLabourDay(),
		core.WithHolidayFunc(variableHolidays),
	)
)

func variableHolidays(date time.Time) core.Holidays {
	if date.Year() > 2018 {
		return core.Holidays{
			fmt.Sprintf("%d/5/9", date.Year()): core.Event("Europe Day"),
		}
	}
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
