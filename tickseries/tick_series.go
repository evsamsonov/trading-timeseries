package tickseries

import (
	"errors"
)

var (
	ErrCannotBeNil  = errors.New("tick cannot be nil")
	ErrAlreadyExist = errors.New("tick is already exist")
	ErrEarlierTime  = errors.New("tick time is earlier than previous")
)

// TickSeries represents TickSeries of trading ticks
type TickSeries struct {
	ticks       []*Tick
	ids         map[int64]struct{}
	allowZeroID bool
}

type Option func(*TickSeries)

func WithAllowZeroIDOption(allowZeroID bool) Option {
	return func(ts *TickSeries) {
		ts.allowZeroID = allowZeroID
	}
}

// New creates and returns new TickSeries
func New(opts ...Option) *TickSeries {
	ts := &TickSeries{
		ticks: make([]*Tick, 0),
		ids:   make(map[int64]struct{}),
	}
	for _, opt := range opts {
		opt(ts)
	}
	return ts
}

// Add adds a trading tick to the TickSeries. It allows to add
// a tick with unique ID and with later or equal time than
// last added tick
func (t *TickSeries) Add(tick *Tick) error {
	if tick == nil {
		return ErrCannotBeNil
	}

	if !t.allowZeroID {
		if _, ok := t.ids[tick.ID]; ok {
			return ErrAlreadyExist
		}
	}

	last := t.Last()
	if last != nil && tick.Time.Before(last.Time) {
		return ErrEarlierTime
	}

	t.ticks = append(t.ticks, tick)
	t.ids[tick.ID] = struct{}{}

	return nil
}

// Tick returns tick by index
func (t *TickSeries) Tick(i int) *Tick {
	if i >= 0 && i < len(t.ticks) {
		return t.ticks[i]
	}

	return nil
}

// Last returns last tick in TickSeries
func (t *TickSeries) Last() *Tick {
	if len(t.ticks) > 0 {
		return t.ticks[len(t.ticks)-1]
	}

	return nil
}

// Length returns length of TickSeries
func (t *TickSeries) Length() int {
	return len(t.ticks)
}

// Iterator returns channel for iterate by TickSeries. It requires
// to get all messages from channel to avoid goroutine leak
func (t *TickSeries) Iterator() chan *Tick {
	ch := make(chan *Tick)
	go func() {
		for _, tick := range t.ticks {
			ch <- tick
		}
		close(ch)
	}()

	return ch
}
