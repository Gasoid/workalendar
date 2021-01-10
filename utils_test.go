package workalendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWesternEaster(t *testing.T) {
	easter2019 := time.Date(2019, time.April, 21, 0, 0, 0, 0, time.UTC)
	easter2020 := time.Date(2020, time.April, 12, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, easter2019, easter(2019, EasterWestern))
	assert.Equal(t, easter2020, easter(2020, EasterWestern))
	assert.NotEqual(t, easter2020, easter(2019, EasterWestern))
}

func TestOrthodoxEaster(t *testing.T) {
	easter2019 := time.Date(2019, time.April, 28, 0, 0, 0, 0, time.UTC)
	easter2020 := time.Date(2020, time.April, 19, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, easter2019, easter(2019, EasterOrthodox))
	assert.Equal(t, easter2020, easter(2020, EasterOrthodox))
	assert.NotEqual(t, easter2020, easter(2019, EasterOrthodox))
}

func TestFindWorkingDay(t *testing.T) {
	jan10 := time.Date(2021, time.January, 10, 0, 0, 0, 0, time.UTC)
	monday := jan10.AddDate(0, 0, 1)
	jan9 := jan10.AddDate(0, 0, -1)
	friday := time.Date(2021, time.January, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, monday, FindWorkingDay(jan10))
	assert.Equal(t, monday, FindWorkingDay(jan9))
	assert.Equal(t, friday, FindWorkingDay(friday))
}

func TestFirstMonday(t *testing.T) {
	jan1 := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	jan4 := time.Date(2021, time.January, 4, 0, 0, 0, 0, time.UTC)
	feb1 := time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, jan4, FindFirstMonday(jan1))
	assert.Equal(t, feb1, FindFirstMonday(feb1))
}

func TestLastMonday(t *testing.T) {
	jan4 := time.Date(2021, time.January, 4, 0, 0, 0, 0, time.UTC)
	jan25 := time.Date(2021, time.January, 25, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, jan25, FindLastMonday(jan4))
}
