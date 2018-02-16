package event

import (
	"container/heap"
	"fmt"
)

// Queue is a priority queue of events ordered by scheduled time and event ID
type Queue interface {
	// SheduleEvent schedules an event for execution
	ScheduleEvent(at uint64, event Event)

	// Pop remove and returns the next event to be executed from the queue.
	Pop() Event

	// Len returns the current size of the queue
	Len() int
}

type queue struct {
	heap   Heap
	nextID uint64
}

func (q *queue) Pop() Event {
	wrappedEvent := heap.Pop(&q.heap).(scheduledEvent)
	return wrappedEvent.event
}

func (q *queue) ScheduleEvent(when uint64, event Event) {
	fmt.Println("Next event ID:", q.nextID)
	wrappedEvent := scheduledEvent{
		scheduledAt: when,
		id:          q.nextID,
		event:       event}

	heap.Push(&q.heap, wrappedEvent)
	q.nextID++
}

func (q *queue) Len() int {
	return q.heap.Len()
}

// NewQueue returns a new event queue
func NewQueue() Queue {
	q := &queue{
		heap:   Heap{},
		nextID: 0}

	heap.Init(&q.heap)

	return q
}
