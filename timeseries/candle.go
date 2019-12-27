package timeseries

import "time"

// Candle represents trading candle
type Candle struct {
	Time   time.Time `json:"time"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Open   float64   `json:"open"`
	Close  float64   `json:"close"`
	Volume int64     `json:"volume"`
}

// NewCandle creates new candle
func NewCandle(time time.Time) *Candle {
	return &Candle{Time: time}
}
