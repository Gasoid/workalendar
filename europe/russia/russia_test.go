package russia_test

import (
	"testing"
	"time"

	"github.com/Gasoid/workalendar/europe/russia"
	"github.com/stretchr/testify/assert"
)

var (
	march8     = time.Date(2020, time.March, 8, 0, 0, 0, 0, time.UTC)
	may9       = time.Date(2020, time.May, 11, 0, 0, 0, 0, time.UTC)
	notHoliday = time.Date(2020, time.May, 17, 0, 0, 0, 0, time.UTC)
)

func TestIsHoliday(t *testing.T) {
	assert.True(t, russia.IsHoliday(march8))
	assert.True(t, russia.IsHoliday(may9))
}

func TestGetHoliday(t *testing.T) {
	holiday, err := russia.GetHoliday(may9)
	v := "Victory Day"
	assert.NoError(t, err)
	assert.Equal(t, v, holiday.Name)
}

func TestGetHoliday_NotHoliday(t *testing.T) {
	_, err := russia.GetHoliday(notHoliday)
	assert.Error(t, err)
}

func TestIsWorkingDay(t *testing.T) {
	assert.False(t, russia.IsWorkingDay(march8))
	assert.False(t, russia.IsWorkingDay(may9))
}
