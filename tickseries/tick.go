package tickseries

import "time"

type Operation int

const (
	Buy Operation = iota + 1
	Sell
)

type Tick struct {
	ID        int64
	Time      time.Time
	Price     float64
	Volume    int64
	Operation Operation
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
