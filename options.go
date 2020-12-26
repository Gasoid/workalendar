package workalendar

import "time"

const (
	newYearDay           string = "1/1"
	labourDay            string = "5/1"
	christmasDay         string = "12/25"
	christmasEve         string = "12/24"
	boxingDay            string = "12/26"
	ortodoxChristmasDay  string = "1/7"
	epiphanyDay          string = "1/6"
	assumptionDay        string = "8/15"
	allSaints            string = "11/1"
	immaculateConception string = "12/8"
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
		c.Days[ortodoxChristmasDay] = Event("Christmas")
	}
}

//WithChristmas is option implementing christmas
func WithChristmas() CalendarOption {
	return func(c *Calendar) {
		c.Days[christmasDay] = Event("Christmas Day")
	}
}

//WithChristmasEve is option implementing christmas eve
func WithChristmasEve() CalendarOption {
	return func(c *Calendar) {
		c.Days[christmasEve] = Event("Christmas Eve")
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

//WithEasterSaturday is option implementing easter Saturday
func WithEasterSaturday() CalendarOption {
	return func(c *Calendar) {
		c.easterMethod = EasterWestern
		c.includeEasterSaturday = true
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

//WithEasterTuesday is option implementing easter tuesday
func WithEasterTuesday() CalendarOption {
	return func(c *Calendar) {
		c.includeEasterTuesday = true
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

//WithEpiphany is option implementing good friday
func WithEpiphany() CalendarOption {
	return func(c *Calendar) {
		c.Days[epiphanyDay] = Event("Epiphany")
	}
}

//WithCorpusChristi is option implementing Corpus Christi
func WithCorpusChristi() CalendarOption {
	return func(c *Calendar) {
		c.includeCorpusChristi = true
	}
}

//WithAssumption is option implementing Assumption
func WithAssumption() CalendarOption {
	return func(c *Calendar) {
		c.Days[assumptionDay] = Event("Assumption of Mary to Heaven")
	}
}

//WithAllSaints is option implementing All Saints
func WithAllSaints() CalendarOption {
	return func(c *Calendar) {
		c.Days[allSaints] = Event("All Saints Day")
	}
}

//WithImmaculateConception is option implementing Immaculate Conception
func WithImmaculateConception() CalendarOption {
	return func(c *Calendar) {
		c.Days[immaculateConception] = Event("Immaculate Conception")
	}
}

//WithRadonitsa is option implementing Radonitsa
func WithRadonitsa() CalendarOption {
	return func(c *Calendar) {
		c.includeRadonitsa = true
	}
}

//WithCleanMonday is option implementing clean monday
func WithCleanMonday() CalendarOption {
	return func(c *Calendar) {
		c.includeCleanMonday = true
	}
}

//WithHolyThursday is option implementing holy thursday
func WithHolyThursday() CalendarOption {
	return func(c *Calendar) {
		c.includeHolyThursday = true
	}
}
