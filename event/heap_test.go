package event

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapInitialLength(t *testing.T) {

	h := &Heap{}
	heap.Init(h)

	if h.Len() != 0 {
		t.Errorf("Heap should be empty")
	}
}

func newScheduledEvent(at uint64, id uint64) scheduledEvent {
	return scheduledEvent{
		scheduledAt: at,
		id:          id,
		event:       NewNoopEvent()}
}

func TestHeapPushPop(t *testing.T) {

	h := &Heap{}
	heap.Init(h)

	assert.Equal(t,
		0, h.Len(),
		"Heap should be empty")

	eventA := newScheduledEvent(100, 1)
	eventB := newScheduledEvent(100, 2)
	eventC := newScheduledEvent(10, 3)
	eventD := newScheduledEvent(99, 4)

	heap.Push(h, eventA)
	heap.Push(h, eventB)
	heap.Push(h, eventC)

	assert.Equal(t,
		3, h.Len(),
		"Heap length should be 2, but was %d", h.Len())

	assert.Equal(t, eventC, heap.Pop(h))

	// eventC schedules eventD ahead of eventA and eventB
	heap.Push(h, eventD)

	assert.Equal(t, eventD, heap.Pop(h))
	assert.Equal(t, eventA, heap.Pop(h))
	assert.Equal(t, eventB, heap.Pop(h))

	assert.Equal(t,
		0, h.Len(),
		"Heap should be empty")
}
