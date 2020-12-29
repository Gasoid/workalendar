package workalendar

import (
	"fmt"
	"time"
)

//CalEvent describes holiday
type CalEvent struct {
	Name string
}

func (e *CalEvent) String() string {
	return e.Name
}

//Event returns new CalEvent
func Event(name string) *CalEvent {
	return &CalEvent{Name: name}
}

//CalendarOption is inteded to be passed to NewCalendar
type CalendarOption func(*Calendar)

//NewCalendar prepares calendar struct
func NewCalendar(holidays Holidays, opts ...CalendarOption) *Calendar {
	c := &Calendar{Days: holidays}
	for _, opt := range opts {
		opt(c)
	}
	if c.easterMethod == 0 {
		c.easterMethod = EasterWestern
	}
	return c
}

//Holidays is a map to represent events
type Holidays map[string]*CalEvent

//Calendar is a struct that is intended for representing all holidays of country
type Calendar struct {
	Days               Holidays
	additionalHolidays func(date time.Time) Holidays
	easterMethod       easterType
	includeEasterSaturday,
	includeEasterSunday,
	includeEasterMonday,
	includeEasterTuesday,
	includeGoodFriday,
	includeAscension,
	includeCorpusChristi,
	includeWhitMonday,
	includeWhitSunday,
	includeRadonitsa,
	includeCleanMonday,
	includeHolyThursday bool
}

//CheckHoliday is intended to determine whether day is holiday
func (c *Calendar) CheckHoliday(date time.Time) (bool, *CalEvent) {
	year, month, day := date.Date()
	if event, ok := c.Days[fmt.Sprintf("%d/%d", month, day)]; ok {
		return true, event
	}
	if event, ok := c.Days[fmt.Sprintf("%d/%d/%d", year, month, day)]; ok {
		return true, event
	}
	days := c.additionalHolidays(date)
	if event, ok := days[fmt.Sprintf("%d/%d/%d", year, month, day)]; ok {
		return true, event
	}
	if ok, event := c.checkEasterHolidays(date); ok {
		return true, event
	}
	if ok, event := c.checkOrthodoxEasterHolidays(date); ok {
		return true, event
	}
	return false, nil
}

func (c *Calendar) checkOrthodoxEasterHolidays(date time.Time) (bool, *CalEvent) {
	easterSunday := easter(date.Year(), EasterOrthodox)
	radonitsa := easterSunday.AddDate(0, 0, 9)
	if c.includeRadonitsa {
		if radonitsa.Month() == date.Month() && radonitsa.Day() == date.Day() {
			return true, Event("Radonitsa")
		}
	}
	return false, nil
}

func (c *Calendar) checkEasterHolidays(date time.Time) (bool, *CalEvent) {
	easterSunday := easter(date.Year(), c.easterMethod)
	if c.includeEasterSunday {
		if easterSunday.Month() == date.Month() && easterSunday.Day() == date.Day() {
			return true, Event("Easter Sunday")
		}
	}
	if c.includeEasterSaturday {
		easterSaturday := easterSunday.AddDate(0, 0, -1)
		if easterSaturday.Month() == date.Month() && easterSaturday.Day() == date.Day() {
			return true, Event("Easter Saturday")
		}
	}
	if c.includeEasterMonday {
		easterMonday := easterSunday.AddDate(0, 0, 1)
		if easterMonday.Month() == date.Month() && easterMonday.Day() == date.Day() {
			return true, Event("Easter Monday")
		}
	}
	if c.includeEasterTuesday {
		easterTuesday := easterSunday.AddDate(0, 0, 2)
		if easterTuesday.Month() == date.Month() && easterTuesday.Day() == date.Day() {
			return true, Event("Easter Tuesday")
		}
	}
	if c.includeGoodFriday {
		goodFriday := easterSunday.AddDate(0, 0, -2)
		if goodFriday.Month() == date.Month() && goodFriday.Day() == date.Day() {
			return true, Event("Good Friday")
		}
	}
	if c.includeAscension {
		ascensionThursday := easterSunday.AddDate(0, 0, 39)
		if ascensionThursday.Month() == date.Month() && ascensionThursday.Day() == date.Day() {
			return true, Event("Ascension Thursday")
		}
	}
	//TODO: make sure whether it is right
	if c.includeCorpusChristi {
		corpusChristi := easterSunday.AddDate(0, 0, 60)
		if corpusChristi.Month() == date.Month() && corpusChristi.Day() == date.Day() {
			return true, Event("Corpus Christi")
		}
	}
	if c.includeCleanMonday {
		cleanMonday := easterSunday.AddDate(0, 0, 48)
		if cleanMonday.Month() == date.Month() && cleanMonday.Day() == date.Day() {
			return true, Event("Clean Monday")
		}
	}
	if c.includeHolyThursday {
		holyThursday := easterSunday.AddDate(0, 0, 26)
		if holyThursday.Month() == date.Month() && holyThursday.Day() == date.Day() {
			return true, Event("Holy Thursday")
		}
	}
	return false, nil
}

func (c *Calendar) checkWhitHolidays(date time.Time) (bool, *CalEvent) {
	easterSunday := easter(date.Year(), c.easterMethod)
	whitMonday := easterSunday.AddDate(0, 0, 50)
	whitSunday := easterSunday.AddDate(0, 0, 49)
	if c.includeWhitMonday {
		if whitMonday.Month() == date.Month() && whitMonday.Day() == date.Day() {
			return true, Event("Whit Monday")
		}
	}
	if c.includeWhitSunday {
		if whitSunday.Month() == date.Month() && whitSunday.Day() == date.Day() {
			return true, Event("Whit Sunday")
		}
	}
	return false, nil
}

// func (c *Calendar) checkChristianHolidays(date time.Time) (bool, *CalEvent) {
// 	return false, nil
// }

//IsWorkingDay is inteded to check whether a day is working or not
func (c *Calendar) IsWorkingDay(date time.Time) bool {
	if date.Weekday() == time.Sunday || date.Weekday() == time.Saturday {
		return false
	}
	ok, _ := c.CheckHoliday(date)
	return !ok
}

//IsHoliday is inteded to check whether a day is holiday or not
func (c *Calendar) IsHoliday(date time.Time) bool {
	ok, _ := c.CheckHoliday(date)
	return ok
}

//GetHoliday is inteded to check whether a day is holiday or not
func (c *Calendar) GetHoliday(date time.Time) *CalEvent {
	ok, event := c.CheckHoliday(date)
	if ok {
		return event
	}
	return nil
}
