package timeseries

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCandle(t *testing.T) {
	now := time.Now()
	got := NewCandle(now)

	assert.Equal(t, &Candle{Time: now}, got)
}
