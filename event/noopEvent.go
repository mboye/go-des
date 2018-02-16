package event

// NoopEvent does nothing when executed
type NoopEvent struct {
	scheduledAt uint64
	id          uint64
}

// ScheduledAt returns when the event will be executed
func (e NoopEvent) ScheduledAt() uint64 {
	return e.scheduledAt
}

// ID returns the event ID
func (e NoopEvent) ID() uint64 {
	return e.id
}

// Execute executes a NoopEvent
func (e NoopEvent) Execute() {
}

// NewNoopEvent return a new NoopEvent
func NewNoopEvent(scheduledAt, id uint64) *NoopEvent {
	return &NoopEvent{
		scheduledAt: scheduledAt,
		id:          id}
}
