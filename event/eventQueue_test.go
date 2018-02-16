package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueInit(t *testing.T) {
	q := NewQueue()

	if q.Len() != 0 {
		t.Errorf("Queue should be empty")
	}
}

func TestQueueScheduleEvent(t *testing.T) {
	q := NewQueue()

	eventA := NewNoopEvent()
	eventB := NewNoopEvent()

	q.ScheduleEvent(eventA, 10)
	q.ScheduleEvent(eventB, 10)

	assert.Equal(t, 2, q.Len(), "Queue length should be 2, but was %d", q.Len())

	assert.Equal(t, eventA, q.Pop(), "First event should be eventA")
	assert.Equal(t, eventB, q.Pop(), "Second event should be eventB")
}

func TestQueueClock(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, 0, q.Clock())

	q.ScheduleEvent(NewNoopEvent(), 10)
	assert.Equal(t, 0, q.Clock())
	q.Pop()
	assert.Equal(t, 10, q.Clock())
}
