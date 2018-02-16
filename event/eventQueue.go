package event

import (
	"container/heap"
	"fmt"
)

// Queue is a priority queue of events ordered by scheduled time and event ID
type Queue interface {
	EventScheduler

	// Pop remove and returns the next event to be executed from the queue.
	Pop() Event

	// Len returns the current size of the queue
	Len() int

	// Clock returns the time of the last popped event
	Clock() int
}

type queue struct {
	heap   Heap
	nextID int
	clock  int
}

func (q *queue) updateClock(newClock int) {
	if newClock > q.clock {
		q.clock = newClock
	}
}

func (q *queue) Pop() Event {
	if q.heap.Len() == 0 {
		return NewSimulationStopEvent("End of simulation reached")
	}

	scheduledEvent := heap.Pop(&q.heap).(scheduledEvent)
	q.updateClock(scheduledEvent.scheduledAt)
	return scheduledEvent.event
}

func (q *queue) ScheduleEvent(event Event, delay int) {
	fmt.Println("Next event ID:", q.nextID)
	wrappedEvent := scheduledEvent{
		scheduledAt: delay,
		id:          q.nextID,
		event:       event}

	heap.Push(&q.heap, wrappedEvent)
	q.nextID++
}

func (q *queue) Len() int {
	return q.heap.Len()
}

func (q *queue) Clock() int {
	return q.clock
}

// NewQueue returns a new event queue
func NewQueue() Queue {
	q := &queue{
		heap:   Heap{},
		nextID: 0}

	heap.Init(&q.heap)

	return q
}
