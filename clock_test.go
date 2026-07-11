package clock

import (
	"testing"
	"time"
)

func TestFakeAdvance(t *testing.T) {
	start := time.Date(2026, 6, 14, 0, 0, 0, 0, time.UTC)
	c := NewFake(start)

	if !c.Now().Equal(start) {
		t.Fatalf("Now() = %v, want %v", c.Now(), start)
	}

	c.Advance(48 * time.Hour)
	want := start.Add(48 * time.Hour)
	if !c.Now().Equal(want) {
		t.Fatalf("after Advance, Now() = %v, want %v", c.Now(), want)
	}
}

func TestSystemNowIsUTC(t *testing.T) {
	if loc := (System{}).Now().Location(); loc != time.UTC {
		t.Fatalf("System.Now() location = %v, want UTC", loc)
	}
}
