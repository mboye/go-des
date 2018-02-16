package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLessComparesByIDWhenScheduledAtSameTime(t *testing.T) {
	eventA := scheduledEvent{100, 1, nil}
	eventB := scheduledEvent{100, 2, nil}

	if Less(eventA, eventB) != true {
		t.Errorf("Expected eventA to be less than eventB")
	}

	if Less(eventB, eventA) != false {
		t.Errorf("Expected eventB not to be less than eventA")
	}
}

func TestLessComparesByScheduledAt(t *testing.T) {
	eventA := scheduledEvent{100, 1, nil}
	eventB := scheduledEvent{101, 2, nil}

	if Less(eventA, eventB) != true {
		t.Errorf("Expected eventA to be less than eventB")
	}

	if Less(eventB, eventA) != false {
		t.Errorf("Expected eventB not to be less than eventA")
	}
}

func TestLessPanicsWhenIDsEqual(t *testing.T) {
	eventA := scheduledEvent{100, 1, nil}
	eventB := scheduledEvent{100, 1, nil}

	assert.Panics(
		t,
		func() { Less(eventA, eventB) },
		"Less() should panic if event IDs are the same")
}
