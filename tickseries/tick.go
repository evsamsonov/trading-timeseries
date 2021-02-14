package tickseries

import "time"

type Operation int

const (
	Buy Operation = iota + 1
	Sell
)

type Tick struct {
	ID        int64     `json:"id" validate:"required"`
	Time      time.Time `json:"time" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
	Volume    int64     `json:"volume" validate:"required"`
	Operation Operation `json:"operation" validate:"required,oneof=1 2"`
}

func NewTick(ID int64) *Tick {
	return &Tick{
		ID: ID,
	}
}

func (t *Tick) IsBuy() bool {
	return t.Operation == Buy
}

func (t *Tick) IsSell() bool {
	return t.Operation == Sell
}
