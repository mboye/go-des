package event

// Processor executes simulation events
type Processor interface {
	Name() string
	Process(e Event, scheduler EventScheduler)
	Subscriptions() []string
}
