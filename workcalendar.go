package workalendar

import (
	"fmt"
	"time"
)

const (
	christmasDay        string = "12/25"
	ortodoxChristmasDay string = "1/7"
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

//WithHolidayFunc is inteded to add HolidayFunc
func WithHolidayFunc(additionalHolidaysFunc func(date time.Time) Holidays) CalendarOption {
	return func(c *Calendar) {
		c.additionalHolidays = additionalHolidaysFunc
	}
}

//WithOrtodoxChristmas is option implementing ortodox christmas
func WithOrtodoxChristmas() CalendarOption {
	return func(c *Calendar) {
		c.Days[ortodoxChristmasDay] = Event("Christmas Day")
	}
}

//WithChristmas is option implementing christmas
func WithChristmas() CalendarOption {
	return func(c *Calendar) {
		c.Days[christmasDay] = Event("Christmas Day")
	}
}

//NewCalendar prepares calendar struct
func NewCalendar(holidays Holidays, opts ...CalendarOption) *Calendar {
	c := &Calendar{Days: holidays}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

//Holidays is a map to represent events
type Holidays map[string]*CalEvent

//Calendar is a struct that is intended for representing all holidays of country
type Calendar struct {
	Days               Holidays
	additionalHolidays func(date time.Time) Holidays
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
	return false, nil
}

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
