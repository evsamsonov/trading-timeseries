package timeseries

import (
	"fmt"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestNewTimeSeries(t *testing.T) {
	got := New()

	assert.Equal(t, &TimeSeries{[]*Candle{}}, got)
}

func TestTimeSeries_AddCandle(t *testing.T) {
	timeSeries := New()

	t.Run("candle=nil", func(t *testing.T) {
		err := timeSeries.AddCandle(nil)
		assert.EqualError(t, err, "candle cannot be nil")
	})

	t.Run("open=0", func(t *testing.T) {
		candle := createTestCandle()
		candle.Open = 0
		err := timeSeries.AddCandle(candle)

		assert.EqualError(t, err, "open cannot be 0")
	})

	t.Run("close=0", func(t *testing.T) {
		candle := createTestCandle()
		candle.Close = 0
		err := timeSeries.AddCandle(candle)

		assert.EqualError(t, err, "close cannot be 0")
	})

	t.Run("high=0", func(t *testing.T) {
		candle := createTestCandle()
		candle.High = 0
		err := timeSeries.AddCandle(candle)

		assert.EqualError(t, err, "high cannot be 0")
	})

	t.Run("low=0", func(t *testing.T) {
		candle := createTestCandle()
		candle.Low = 0
		err := timeSeries.AddCandle(candle)

		assert.EqualError(t, err, "low cannot be 0")
	})

	t.Run("LastCandle=nil", func(t *testing.T) {
		candle := createTestCandle()
		err := timeSeries.AddCandle(candle)

		assert.Nil(t, err)
		assert.Equal(t, []*Candle{candle}, timeSeries.candles)
	})

	t.Run("Time after Last candle time", func(t *testing.T) {
		timeSeries := New()

		candle1 := createTestCandle()
		candle2 := createTestCandle()
		candle2.Time = time.Unix(2, 0)

		err1 := timeSeries.AddCandle(candle1)
		err2 := timeSeries.AddCandle(candle2)

		assert.Nil(t, err1)
		assert.Nil(t, err2)
		assert.Equal(t, []*Candle{candle1, candle2}, timeSeries.candles)
	})

	t.Run("Time before Last candle time", func(t *testing.T) {
		timeSeries := New()

		candle1 := createTestCandle()
		candle2 := createTestCandle()
		candle2.Time = time.Unix(1, 0)

		err1 := timeSeries.AddCandle(candle1)
		err2 := timeSeries.AddCandle(candle2)

		assert.Nil(t, err1)
		assert.EqualError(t, err2, fmt.Sprintf("time is earlier or equal previous"), )
		assert.Equal(t, []*Candle{candle1}, timeSeries.candles)
	})
}

func TestTimeSeries_LastCandle(t *testing.T) {
	t.Run("LastCandle=nil", func(t *testing.T) {
		series := New()
		assert.Nil(t, series.LastCandle())
	})

	t.Run("LastCandle not nil", func(t *testing.T) {
		series := New()

		candle := createTestCandle()
		candle.Time = time.Unix(2, 0)

		err1 := series.AddCandle(createTestCandle())
		err2 := series.AddCandle(candle)

		assert.Nil(t, err1)
		assert.Nil(t, err2)
		assert.Equal(t, candle, series.LastCandle())
	})
}

func TestTimeSeries_Candle(t *testing.T) {
	t.Run("Out of range", func(t *testing.T) {
		series := New()
		err := series.AddCandle(createTestCandle())

		assert.Nil(t, err)
		assert.Nil(t, series.Candle(-1))
		assert.Nil(t, series.Candle(1))
	})

	t.Run("Out of range", func(t *testing.T) {
		series := New()

		candle1 := createTestCandle()
		candle2 := createTestCandle()
		candle2.Time = time.Unix(2, 0)

		err1 := series.AddCandle(candle1)
		err2 := series.AddCandle(candle2)

		assert.Nil(t, err1)
		assert.Nil(t, err2)
		assert.Equal(t, candle1, series.Candle(0))
		assert.Equal(t, candle2, series.Candle(1))
	})
}

func createTestCandle() *Candle {
	return &Candle{Open: 1, Close: 1, High: 1, Low: 1, Time: time.Unix(1, 0)}
}

func TestTimeSeries_Length(t *testing.T) {
	series := New()

	assert.Nil(t, series.AddCandle(createTestCandle()))

	candle2 := createTestCandle()
	candle2.Time = time.Unix(2, 0)
	assert.Nil(t, series.AddCandle(candle2))

	assert.Equal(t, 2, series.Length())
}
