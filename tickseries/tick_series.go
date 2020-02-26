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

type TickSeries struct {
	ticks []*Tick
	ids   map[int64]struct{}
	mu    sync.RWMutex
}

func NewTickSeries() *TickSeries {
	return &TickSeries{
		ticks: make([]*Tick, 0),
		ids:   make(map[int64]struct{}),
	}
}

func (t *TickSeries) Add(tick *Tick) error {
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

func (t *TickSeries) Tick(i int) *Tick {
	if i >= 0 && i < len(t.ticks) {
		return t.ticks[i]
	}

	return nil
}

func (t *TickSeries) Last() *Tick {
	if len(t.ticks) > 0 {
		return t.ticks[len(t.ticks)-1]
	}

	return nil
}

func (t *TickSeries) Length() int {
	return len(t.ticks)
}

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
