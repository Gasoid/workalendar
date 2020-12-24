package workalendar

import "time"

const (
	newYearDay          string = "1/1"
	labourDay           string = "5/1"
	christmasDay        string = "12/25"
	boxingDay           string = "12/26"
	ortodoxChristmasDay string = "1/7"
)

//WithHolidayFunc is inteded to add HolidayFunc
func WithHolidayFunc(additionalHolidaysFunc func(date time.Time) Holidays) CalendarOption {
	return func(c *Calendar) {
		c.additionalHolidays = additionalHolidaysFunc
	}
}

//WithOrthodoxChristmas is option implementing orthodox christmas
func WithOrthodoxChristmas() CalendarOption {
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

//WithBoxingDay is option implementing boxing day
func WithBoxingDay() CalendarOption {
	return func(c *Calendar) {
		c.Days[boxingDay] = Event("Boxing Day")
	}
}

//WithNewYearDay is option implementing newyear day
func WithNewYearDay() CalendarOption {
	return func(c *Calendar) {
		c.Days[newYearDay] = Event("New Year")
	}
}

//WithLabourDay is option implementing labour day
func WithLabourDay() CalendarOption {
	return func(c *Calendar) {
		c.Days[labourDay] = Event("Labour Day")
	}
}

//WithEasterSunday is option implementing easter sunday
func WithEasterSunday() CalendarOption {
	return func(c *Calendar) {
		c.easterMethod = EasterWestern
		c.includeEasterSunday = true
	}
}

//WithEasterMonday is option implementing easter monday
func WithEasterMonday() CalendarOption {
	return func(c *Calendar) {
		c.includeEasterMonday = true
	}
}

//WithWhitMonday is option implementing whit monday
func WithWhitMonday() CalendarOption {
	return func(c *Calendar) {
		c.includeWhitMonday = true
	}
}

//WithWhitSunday is option implementing whit sunday
func WithWhitSunday() CalendarOption {
	return func(c *Calendar) {
		c.includeWhitSunday = true
	}
}

//WithGoodFriday is option implementing good friday
func WithGoodFriday() CalendarOption {
	return func(c *Calendar) {
		c.includeGoodFriday = true
	}
}

//WithAscension is option implementing good friday
func WithAscension() CalendarOption {
	return func(c *Calendar) {
		c.includeAscension = true
	}
}
