package belarus

import (
	core "github.com/Gasoid/workalendar"
)

var (
	holidays = core.Holidays{
		"1/2":  core.Event("Day After New Year"),
		"3/8":  core.Event("International Women's Day"),
		"5/9":  core.Event("Victory Day"),
		"7/3":  core.Event("Republic Day"),
		"11/7": core.Event("October Revolution Day"),
	}
	calendar = core.NewCalendar(
		holidays,
		core.WithLabourDay(),
		core.WithOrthodoxChristmas(),
		core.WithChristmas(),
		core.WithRadonitsa(),
	)
)
