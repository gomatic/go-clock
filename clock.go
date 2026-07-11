// Package clock provides an injectable time source so time-dependent
// behavior (alias expiry, audit timestamps) can be tested deterministically.
package clock

import (
	"sync"
	"time"
)

// Clock is the minimal time source consumed by services.
type Clock interface {
	Now() time.Time
}

// System is the production clock; it returns the current UTC time.
type System struct{}

func (System) Now() time.Time { return time.Now().UTC() }

// Fake is a controllable clock for tests.
type Fake struct {
	now time.Time
	mu  sync.Mutex
}

// NewFake returns a Fake anchored at the given instant (coerced to UTC).
func NewFake(t time.Time) *Fake { return &Fake{now: t.UTC()} }

func (f *Fake) Now() time.Time {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.now
}

// Advance moves the clock forward by d.
func (f *Fake) Advance(d time.Duration) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.now = f.now.Add(d)
}
