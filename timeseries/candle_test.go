package timeseries

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewCandle(t *testing.T) {
	now := time.Now()
	got := NewCandle(now)

	assert.Equal(t, &Candle{Time: now}, got)
}
