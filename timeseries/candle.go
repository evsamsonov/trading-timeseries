package timeseries

import "time"

// Candle represents trading candle
type Candle struct {
	Time       time.Time `json:"time" db:"time" validate:"required"`
	High       float64   `json:"high" db:"high" validate:"required"`
	Low        float64   `json:"low" db:"low" validate:"required"`
	Open       float64   `json:"open" db:"open" validate:"required"`
	Close      float64   `json:"close" db:"close" validate:"required"`
	Volume     int64     `json:"volume" db:"volume" validate:"required"`
	IsComplete bool      `json:"is_complete" db:"is_complete"`
}

// NewCandle creates new candle
func NewCandle(time time.Time) *Candle {
	return &Candle{Time: time}
}

// Copy returns a candle copy
func (c *Candle) Copy() Candle {
	return *c
}
