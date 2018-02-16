package event

// Event represents an event in a simulation
type Event interface {
	Execute()
}

type event struct {
}
