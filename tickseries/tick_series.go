package tickseries

import (
	"errors"
	"sync"
)

var (
	ErrCannotBeNil  = errors.New("tick cannot be nil")
	ErrAlreadyExist = errors.New("tick is already exist")
	ErrEarlierTime  = errors.New("tick time is earlier than previous")
)

// series represents series of trading ticks
type series struct {
	ticks []*Tick
	ids   map[int64]struct{}
	mu    sync.RWMutex
}

// New creates and returns new series
func New() *series {
	return &series{
		ticks: make([]*Tick, 0),
		ids:   make(map[int64]struct{}),
	}
}

// Add adds a trading tick to the series. It allows to add
// a tick with unique ID and with later or equal time than
// last added tick
func (t *series) Add(tick *Tick) error {
	if tick == nil {
		return ErrCannotBeNil
	}

	t.mu.RLock()
	if _, ok := t.ids[tick.ID]; ok {
		t.mu.RUnlock()
		return ErrAlreadyExist
	}
	t.mu.RUnlock()

	last := t.Last()
	if last != nil && tick.Time.Before(last.Time) {
		return ErrEarlierTime
	}

	t.ticks = append(t.ticks, tick)
	t.mu.Lock()
	t.ids[tick.ID] = struct{}{}
	t.mu.Unlock()

	return nil
}

// Tick returns tick by index
func (t *series) Tick(i int) *Tick {
	if i >= 0 && i < len(t.ticks) {
		return t.ticks[i]
	}

	return nil
}

// Last returns last tick in series
func (t *series) Last() *Tick {
	if len(t.ticks) > 0 {
		return t.ticks[len(t.ticks)-1]
	}

	return nil
}

// Length returns length of series
func (t *series) Length() int {
	return len(t.ticks)
}

// Iterator returns channel for iterate by series. It requires
// to get all messages from channel to avoid goroutine leak
func (t *series) Iterator() chan *Tick {
	ch := make(chan *Tick)
	go func() {
		for _, tick := range t.ticks {
			ch <- tick
		}
		close(ch)
	}()

	return ch
}
