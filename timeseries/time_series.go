package timeseries

import (
	"errors"
	"fmt"
)

var ErrUnexpectedTime = errors.New("time is earlier or equal previous")

// TimeSeries represents series of trading candles
type TimeSeries struct {
	candles []*Candle
}

// New creates TimeSeries
func New() *TimeSeries {
	ts := new(TimeSeries)
	ts.candles = make([]*Candle, 0)

	return ts
}

// AddCandle adds trading candle to series.
// Each new candle must have later time than previous
func (ts *TimeSeries) AddCandle(c *Candle) error {
	if c == nil {
		return fmt.Errorf("candle cannot be nil")
	}

	if ts.LastCandle() == nil || c.Time.After(ts.LastCandle().Time) {
		ts.candles = append(ts.candles, c)
		return nil
	}

	return ErrUnexpectedTime
}

// Trim returns selected section of candles from series
// [startIndex] is mandatory
// [endIndex] is optional. If would not be in use, assign it as [endIndex=0]
func (ts *TimeSeries) Trim(startIndex int, endIndex int) (*TimeSeries, error) {
	if ts.Length() == 0 {
		return nil, fmt.Errorf("timeseries cannot be empty")
	}

	if startIndex < 0 {
		return nil, fmt.Errorf("startIndex cannot be negative")
	}

	if endIndex < 0 {
		return nil, fmt.Errorf("endIndex cannot be negative")
	}

	if endIndex != 0 && endIndex == startIndex {
		return nil, fmt.Errorf("endIndex should be greater than startIndex")
	}

	if len(ts.candles) < endIndex {
		return nil, fmt.Errorf("endIndex should be less than equal to candle size")
	}

	newTS := *ts

	if endIndex == 0 {
		endIndex = ts.Length()
	}

	newTS.candles = newTS.candles[startIndex:endIndex]

	return &newTS, nil
}

// LastCandle returns last candle in series or nil if series is empty
func (ts *TimeSeries) LastCandle() *Candle {
	if len(ts.candles) > 0 {
		return ts.candles[len(ts.candles)-1]
	}

	return nil
}

// LastCompleteCandle returns last complete candle and index in series or nil if there are no complete candles
func (ts *TimeSeries) LastCompleteCandle() (*Candle, int) {
	for i := ts.Length() - 1; i >= 0; i-- {
		candle := ts.Candle(i)
		if candle.IsComplete {
			return candle, i
		}
	}

	return nil, 0
}

// Candle returns candle by index [0, len(series)-1] or nil if index out of series
func (ts *TimeSeries) Candle(index int) *Candle {
	if index >= 0 && index < len(ts.candles) {
		return ts.candles[index]
	}

	return nil
}

// Length returns length of series
func (ts *TimeSeries) Length() int {
	return len(ts.candles)
}
