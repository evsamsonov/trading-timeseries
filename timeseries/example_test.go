package timeseries

import (
	"fmt"
	"log"
	"time"
)

func ExampleTimeSeries() {
	dataset := []struct {
		Time   time.Time
		High   float64
		Low    float64
		Open   float64
		Close  float64
		Volume int64
	}{
		{Time: time.Unix(1, 0), High: 1, Low: 2, Open: 3, Close: 4, Volume: 5},
		{Time: time.Unix(2, 0), High: 6, Low: 7, Open: 8, Close: 9, Volume: 10},
	}

	series := New()
	for _, item := range dataset {
		candle := NewCandle(item.Time)
		candle.Open = item.Open
		candle.Close = item.Close
		candle.High = item.High
		candle.Low = item.Low
		candle.Volume = item.Volume

		err := series.AddCandle(candle)
		if err != nil {
			log.Fatalf("Failed to add candle: %v\n", err)
		}
	}

	fmt.Printf("Candle\t\t= %v\n", series.Candle(0))
	fmt.Printf("Last candle\t= %v\n", series.LastCandle())
	fmt.Printf("Length\t\t= %v\n", series.Length())

	series, _ = series.Trim(1, 0)
	fmt.Printf("After trim\t= %v\n", series.Candle(0))
	// Output:
	// Candle		= &{1970-01-01 03:00:01 +0300 MSK 1 2 3 4 5}
	// Last candle	= &{1970-01-01 03:00:02 +0300 MSK 6 7 8 9 10}
	// Length		= 2
	// After trim	= &{1970-01-01 03:00:02 +0300 MSK 6 7 8 9 10}
}
