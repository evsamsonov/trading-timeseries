package timeseries

import "time"

// Candle represents trading candle
type Candle struct {
	Time   time.Time
	High   float64
	Low    float64
	Open   float64
	Close  float64
	Volume int64
}

// NewCandle creates new candle
func NewCandle(time time.Time) *Candle {
	return &Candle{Time: time}
}
