package tickseries

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTickSeries_Add(t *testing.T) {
	t.Run("tick is nil", func(t *testing.T) {
		tickSeries := New()
		assert.Equal(t, ErrCannotBeNil, tickSeries.Add(nil))
	})

	t.Run("tick is already exist", func(t *testing.T) {
		tickSeries := New()
		tick := NewTick(1)
		assert.Equal(t, nil, tickSeries.Add(tick))
		assert.Equal(t, ErrAlreadyExist, tickSeries.Add(tick))
	})

	t.Run("tick time is earlier than previous", func(t *testing.T) {
		tickSeries := New()
		tick1 := NewTick(1)
		tick1.Time = time.Unix(1, 0)
		assert.Equal(t, nil, tickSeries.Add(tick1))

		tick2 := NewTick(2)
		tick2.Time = time.Unix(0, 0)
		assert.Equal(t, ErrEarlierTime, tickSeries.Add(tick2))
	})

	t.Run("added successfully", func(t *testing.T) {
		tickSeries := New()
		tick1 := NewTick(1)
		tick1.Time = time.Unix(1, 0)
		assert.Equal(t, nil, tickSeries.Add(tick1))

		tick2 := NewTick(2)
		tick2.Time = time.Unix(1, 0)
		assert.Equal(t, nil, tickSeries.Add(tick2))

		tick3 := NewTick(3)
		tick3.Time = time.Unix(2, 0)
		assert.Equal(t, nil, tickSeries.Add(tick3))
	})
}

func TestTickSeries_Tick(t *testing.T) {
	tickSeries := New()

	tick := NewTick(1)
	tick.Time = time.Unix(1, 0)
	assert.Equal(t, nil, tickSeries.Add(tick))
	assert.Equal(t, tick, tickSeries.Tick(0))
	assert.Equal(t, (*Tick)(nil), tickSeries.Tick(1))
	assert.Equal(t, (*Tick)(nil), tickSeries.Tick(-1))
}

func TestTickSeries_Last(t *testing.T) {
	tickSeries := New()

	tick1 := NewTick(1)
	tick1.Time = time.Unix(1, 0)
	assert.Equal(t, nil, tickSeries.Add(tick1))

	tick2 := NewTick(2)
	tick2.Time = time.Unix(1, 0)
	assert.Equal(t, nil, tickSeries.Add(tick2))

	assert.Equal(t, tick2, tickSeries.Last())
}

func TestTickSeries_Length(t *testing.T) {
	tickSeries := New()

	tick1 := NewTick(1)
	tick1.Time = time.Unix(1, 0)
	assert.Equal(t, nil, tickSeries.Add(tick1))

	tick2 := NewTick(2)
	tick2.Time = time.Unix(1, 0)
	assert.Equal(t, nil, tickSeries.Add(tick2))

	assert.Equal(t, 2, tickSeries.Length())
}

func TestTickSeries_Iterator(t *testing.T) {
	tickSeries := New()

	tick1 := NewTick(1)
	tick1.Time = time.Unix(1, 0)
	assert.Equal(t, nil, tickSeries.Add(tick1))

	tick2 := NewTick(2)
	tick2.Time = time.Unix(1, 0)
	assert.Equal(t, nil, tickSeries.Add(tick2))

	iterator := tickSeries.Iterator()
	tick := <-iterator
	assert.Equal(t, int64(1), tick.ID)

	tick = <-iterator
	assert.Equal(t, int64(2), tick.ID)

	_, ok := <-iterator
	assert.False(t, ok)
}
