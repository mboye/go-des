package event

// SimulationStopEvent indicates that the last event in a simulation was reached
type simulationStopEvent struct {
	Reason string
}

func (ev *simulationStopEvent) Name() string {
	return "Simulation stop event"
}

// NewSimulationStopEvent returns a simulation stop event
func NewSimulationStopEvent(reason string) Event {
	return &simulationStopEvent{reason}
}
