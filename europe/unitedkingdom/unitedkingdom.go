package unitedkingdom

import (
	"fmt"
	"time"

	core "github.com/Gasoid/workalendar"
)

//TODO: add shiftNewYearsDay
var (
	holidays = core.Holidays{
		"1973/11/14": core.Event("Royal wedding"),
		"1977/6/7":   core.Event("Queen’s Silver Jubilee"),
		"1981/7/29":  core.Event("Royal wedding"),
		"1999/12/31": core.Event("New Year's Eve"),
		"2002/6/3":   core.Event("Queen’s Golden Jubilee"),
		"2011/4/29":  core.Event("Royal Wedding"),
		"2012/6/5":   core.Event("Queen’s Diamond Jubilee"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithGoodFriday(),
		core.WithEasterSunday(),
		core.WithEasterMonday(),
		core.WithBoxingDay(),
		core.WithHolidayFunc(variableHolidays),
	)
)

func variableHolidays(date time.Time) core.Holidays {
	days := core.Holidays{}
	year := date.Year()
	may := time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC)
	if year == 2020 {
		days["2020/5/8"] = core.Event("Early May bank holiday (VE day)")
	} else {
		may = core.FindFirstMonday(may)
		days[may.Format("2006/1/2")] = core.Event("Early May Bank Holiday")
	}
	lastMonday := core.FindLastMonday(may)
	springBankHoliday := lastMonday.Format("2006/1/2")
	switch year {
	case 2012:
		springBankHoliday = "2012/6/4"
	case 1977:
		springBankHoliday = "1977/6/6"
	case 2002:
		springBankHoliday = "2002/6/4"
	}
	days[springBankHoliday] = core.Event("Spring Bank Holiday")
	lastMonday8 := core.FindLastMonday(time.Date(year, 8, 1, 0, 0, 0, 0, time.UTC))
	days[lastMonday8.Format("2006/1/2")] = core.Event("Late Summer Bank Holiday")
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
