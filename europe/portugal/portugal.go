package portugal

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: add Corpus Christi, it is a variable holiday
var (
	holidays = core.Holidays{
		"4/25": core.Event("Dia da Liberdade"),
		"6/10": core.Event("Dia de Portugal"),
		"8/15": core.Event("Assunção de Nossa Senhora"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithGoodFriday(),
		core.WithEasterSunday(),
		core.WithChristmas(),
		core.WithImmaculateConception(),
		core.WithLabourDay(),
		core.WithHolidayFunc(variableHolidays),
	)
)

func variableHolidays(date time.Time) core.Holidays {
	days := core.Holidays{}
	year := date.Year()
	if year > 2015 || year < 2013 {
		days[fmt.Sprintf("%d/10/5", year)] = core.Event("Implantação da República")
		days[fmt.Sprintf("%d/11/1", year)] = core.Event("Todos os santos")
		days[fmt.Sprintf("%d/12/1", year)] = core.Event("Restauração da Independência")
		return days
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
