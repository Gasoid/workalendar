// Code generated github.com/Gasoid/workalendar/generator/gen.go DO NOT EDIT.
package argentina

import (
	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
       "10/26": core.Event("National Holiday"),
       "3/24": core.Event("Día Nacional de la Memoria por la Verdad y la Justicia"),
       "5/25": core.Event("Día de la Revolución de Mayo"),
       "6/20": core.Event("Día Paso a la Inmortalidad del General Manuel Belgrano"),
       "7/9": core.Event("Día de la Independencia"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithNewYearDay(),
		core.WithLabourDay(),
		core.WithFatTuesday(),
		core.WithGoodFriday(),
		core.WithEasterSaturday(),
		core.WithEasterSunday(),
		core.WithChristmas(),
		core.WithImmaculateConception(),
	)
)
