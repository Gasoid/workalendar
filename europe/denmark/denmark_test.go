package denmark_test

import (
	"testing"
	"time"

	country "github.com/Gasoid/workalendar/europe/denmark"
	"github.com/stretchr/testify/assert"
)

var (
	holiday1   = time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC)
	holiday2   = time.Date(2020, time.December, 25, 0, 0, 0, 0, time.UTC)
	notHoliday = time.Date(2020, time.May, 17, 0, 0, 0, 0, time.UTC)
)

func TestIsHoliday(t *testing.T) {
	assert.True(t, country.IsHoliday(holiday1))
	assert.True(t, country.IsHoliday(holiday2))
}

func TestGetHoliday(t *testing.T) {
	holiday, err := country.GetHoliday(holiday1)
	v := "Labour Day"
	assert.NoError(t, err)
	assert.Equal(t, v, holiday.Name)
}

func TestGetHoliday_NotHoliday(t *testing.T) {
	_, err := country.GetHoliday(notHoliday)
	assert.Error(t, err)
}

func TestIsWorkingDay(t *testing.T) {
	assert.False(t, country.IsWorkingDay(holiday1))
	assert.False(t, country.IsWorkingDay(holiday2))
}
