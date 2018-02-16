package event

// EventScheduler schedules an event for execution in the future
type EventScheduler interface {
	ScheduleEvent(event Event, delay int)
}
