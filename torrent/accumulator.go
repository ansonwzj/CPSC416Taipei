package torrent

import (
		"time"
		// "log"
)

// An accumulator that keeps track of the rate of increase.
type Accumulator struct {
	maxRatePeriod time.Duration
	rateSince     time.Time
	last          time.Time
	rate          float64
	total         int64
}

func NewAccumulator(now time.Time, maxRatePeriod time.Duration) (
	acc *Accumulator) {
	acc = &Accumulator{}
	acc.maxRatePeriod = maxRatePeriod
	acc.rateSince = now.Add(time.Second * -1)
	acc.last = acc.rateSince
	acc.rate = 0.0
	acc.total = 0
	return acc
}

func (a *Accumulator) Add(now time.Time, amount int64) {
	// log.Printf("Add    %v   %v  \n", now, amount)
	a.total += amount
	a.rate = (a.rate*float64(a.last.Sub(a.rateSince)) + float64(amount)) /
		float64(now.Sub(a.rateSince))
	a.last = now
	newRateSince := now.Add(-a.maxRatePeriod)
	if a.rateSince.Before(newRateSince) {
		a.rateSince = newRateSince
	}
}

func (a *Accumulator) GetRate(now time.Time) float64 {
	// log.Printf("GetRate   %v   \n", now)
	a.Add(now, 0)
	return a.GetRateNoUpdate()
}

func (a *Accumulator) GetRateNoUpdate() float64 {
	// log.Println("GetRateNoUpdate    ", a)
	return a.rate * float64(time.Second)
}

func (a *Accumulator) DurationUntilRate(now time.Time, newRate float64) time.Duration {
	// log.Printf("Add    %v   %v  \n", now, newRate)
	rate := a.rate
	if rate <= newRate {
		return time.Duration(0)
	}
	dt := float64(now.Sub(a.rateSince))
	return time.Duration(((rate * dt) / newRate) - dt)
}

func (a *Accumulator) getTotal() int64 {
	// log.Println("getTotal    ", a)
	return a.total
}
