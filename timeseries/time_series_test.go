package timeseries

import (
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
		assert.EqualError(t, err2, "time is earlier or equal previous")
		assert.Equal(t, []*Candle{candle1}, timeSeries.candles)
	})
}

func TestTimeSeries_Trim(t *testing.T) {
	timeSeries := New()

	t.Run("candle=nil", func(t *testing.T) {
		_, err := timeSeries.Trim(0, 0)
		assert.EqualError(t, err, "timeseries cannot be empty")
	})

	t.Run("startIndex=negative", func(t *testing.T) {
		err := timeSeries.AddCandle(createTestCandle())
		_, err1 := timeSeries.Trim(-1, 0)

		assert.Nil(t, err)
		assert.EqualError(t, err1, "startIndex cannot be negative")
	})

	t.Run("endIndex=negative", func(t *testing.T) {
		endIndex := -1
		_, err1 := timeSeries.Trim(0, endIndex)

		assert.EqualError(t, err1, "endIndex cannot be negative")
	})

	t.Run("endIndex=negative", func(t *testing.T) {
		endIndex := 1
		_, err1 := timeSeries.Trim(1, endIndex)

		assert.EqualError(t, err1, "endIndex should be greater than startIndex")
	})

	t.Run("endIndex=cannot_bigger_than_timeseries_length", func(t *testing.T) {
		endIndex := 2
		_, err1 := timeSeries.Trim(0, endIndex)

		assert.EqualError(t, err1, "endIndex should be less than equal to candle size")
	})

	t.Run("startIndex=single new ts should return", func(t *testing.T) {
		candle := createTestCandle()
		candle.Time = time.Unix(2, 0)

		err := timeSeries.AddCandle(candle)

		assert.Nil(t, err)
		newTs, _ := timeSeries.Trim(1, 0)

		assert.Greater(t, timeSeries.Length(), newTs.Length())
	})

	t.Run("startIndex&endIndex=single new ts should return", func(t *testing.T) {
		endIndex := 3

		candle1 := createTestCandle()
		candle1.Time = time.Unix(3, 0)
		candle2 := createTestCandle()
		candle2.Time = time.Unix(4, 0)

		err1 := timeSeries.AddCandle(candle1)
		err2 := timeSeries.AddCandle(candle2)

		assert.Nil(t, err1)
		assert.Nil(t, err2)

		newTs, _ := timeSeries.Trim(2, endIndex)

		assert.Greater(t, timeSeries.Length(), newTs.Length())
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
