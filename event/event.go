package event

// Event represents an event in a simulation
type Event interface {
	Name() string
}

type event struct {
}
