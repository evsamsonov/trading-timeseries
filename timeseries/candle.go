package timeseries

import "time"

// Candle represents trading candle
type Candle struct {
	Time   time.Time `json:"time" db:"time"`
	High   float64   `json:"high" db:"high"`
	Low    float64   `json:"low" db:"low"`
	Open   float64   `json:"open" db:"open"`
	Close  float64   `json:"close" db:"close"`
	Volume int64     `json:"volume" db:"volume"`
}

// NewCandle creates new candle
func NewCandle(time time.Time) *Candle {
	return &Candle{Time: time}
}

// Copy returns a candle copy
func (c *Candle) Copy() Candle {
	return *c
}
