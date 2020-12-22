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

//Calendar prepares calendar struct
func Calendar(holidays Holidays, additionalHolidaysFunc func(date time.Time) Holidays) calendar {
	return calendar{
		Days:               holidays,
		additionalHolidays: additionalHolidaysFunc,
	}
}

//Holidays is a map to represent events
type Holidays map[string]*CalEvent

type calendar struct {
	Days               Holidays
	additionalHolidays func(date time.Time) Holidays
}

func (c *calendar) CheckHoliday(date time.Time) (bool, *CalEvent) {
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

func (c *calendar) IsWorkingDay(date time.Time) bool {
	ok, _ := c.CheckHoliday(date)
	return !ok
}

func (c *calendar) IsHoliday(date time.Time) bool {
	ok, _ := c.CheckHoliday(date)
	return ok
}

func (c *calendar) GetHoliday(date time.Time) *CalEvent {
	ok, event := c.CheckHoliday(date)
	if ok {
		return event
	}
	return nil
}
