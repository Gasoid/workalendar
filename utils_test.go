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
