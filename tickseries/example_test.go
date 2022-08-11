package tickseries

import (
	"fmt"
	"log"
	"time"
)

func ExampleTickSeries() {
	dataset := []struct {
		ID        int64
		Time      time.Time
		Price     float64
		Volume    int64
		Operation string
	}{
		{ID: 1, Time: time.Time{}, Price: 123.5, Volume: 10, Operation: "buy"},
		{ID: 2, Time: time.Time{}, Price: 124.1, Volume: 20, Operation: "sell"},
	}

	series := New()
	for _, item := range dataset {
		operation := Buy
		if item.Operation == "sell" {
			operation = Sell
		}
		tick := &Tick{
			ID:        item.ID,
			Time:      item.Time,
			Price:     item.Price,
			Volume:    item.Volume,
			Operation: operation,
		}
		if err := series.Add(tick); err != nil {
			log.Fatalf("Failed to add tick: %v\n", err)
		}
	}

	fmt.Printf("Tick\t\t= %v\n", series.Tick(0))
	fmt.Printf("Last\t\t= %v\n", series.Last())
	fmt.Printf("Length\t= %d\n", series.Length())

	var i int
	for tick := range series.Iterator() {
		fmt.Printf("Tick %d\t\t= %v\n", i, tick)
		i++
	}

	//err = series.Add(tick)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to add tick %d: %w", tick.ID, err)
	//}

	// Output:
	// Tick		= &{1 0001-01-01 00:00:00 +0000 UTC 123.5 10 1}
	// Last		= &{2 0001-01-01 00:00:00 +0000 UTC 124.1 20 2}
	// Length	= 2
	// Tick 0		= &{1 0001-01-01 00:00:00 +0000 UTC 123.5 10 1}
	// Tick 1		= &{2 0001-01-01 00:00:00 +0000 UTC 124.1 20 2}
}
