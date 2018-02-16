package event

// NoopEvent does nothing when executed
type NoopEvent struct {
}

// Execute executes a NoopEvent
func (e NoopEvent) Execute() {
}

// NewNoopEvent return a new NoopEvent
func NewNoopEvent() *NoopEvent {
	return &NoopEvent{}
}
