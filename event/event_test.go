package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLessComparesByIDWhenScheduledAtSameTime(t *testing.T) {
	eventA := NewNoopEvent(100, 1)
	eventB := NewNoopEvent(100, 2)

	if Less(eventA, eventB) != true {
		t.Errorf("Expected eventA to be less than eventB")
	}

	if Less(eventB, eventA) != false {
		t.Errorf("Expected eventB not to be less than eventA")
	}
}

func TestLessComparesByScheduledAt(t *testing.T) {
	eventA := NewNoopEvent(100, 1)
	eventB := NewNoopEvent(101, 2)

	if Less(eventA, eventB) != true {
		t.Errorf("Expected eventA to be less than eventB")
	}

	if Less(eventB, eventA) != false {
		t.Errorf("Expected eventB not to be less than eventA")
	}
}

func TestLessPanicsWhenIDsEqual(t *testing.T) {
	eventA := NewNoopEvent(100, 1)
	eventB := NewNoopEvent(100, 1)

	assert.Panics(
		t,
		func() { Less(eventA, eventB) },
		"Less() should panic if event IDs are the same")
}
