package event

// NoopEvent does nothing when executed
type NoopEvent struct {
}

// Name returns the name of the event
func (e NoopEvent) Name() string {
	return "No operation event"
}

// NewNoopEvent return a new NoopEvent
func NewNoopEvent() *NoopEvent {
	return &NoopEvent{}
}
