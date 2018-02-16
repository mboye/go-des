package simulation

import (
	"github.bus.zalan.do/mboye/des/event"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("simulation")

// Simulation manages event processors and executes a simulation
type Simulation interface {
	// Run simulation and stop when time limit is reached.
	// If the time limit is zero, the simulation will run until
	// there are no events left.
	Run(timeLimit int)

	// AddEventProcessor registers an event processor so that it can receive events
	AddEventProcessor(e event.Processor)

	// Duration returns the duration of the last simulation run
	Duration() int

	// Number of events processed in the last simulation run
	EventCount() int
}

type eventProcessorSlice []event.Processor

type simulation struct {
	queue         event.Queue
	processors    map[string]event.Processor
	subscriptions map[string]eventProcessorSlice
	shouldStop    bool
	eventCount    int
}

// NewSimulation returns a new simulation context
func NewSimulation() Simulation {
	sim := &simulation{
		queue:         event.NewQueue(),
		processors:    make(map[string]event.Processor),
		subscriptions: make(map[string]eventProcessorSlice)}

	sim.AddEventProcessor(sim)
	return sim
}

func (s *simulation) Name() string {
	return "Simulation"
}

func (s *simulation) Subscriptions() []string {
	return []string{"Simulation stop event"}
}

func (s *simulation) Process(ev event.Event, scheduler event.EventScheduler) {
	if ev.Name() == "Simulation stop event" {
		s.shouldStop = true
	}
}

func (s *simulation) AddEventProcessor(processor event.Processor) {
	s.processors[processor.Name()] = processor
	for _, eventName := range processor.Subscriptions() {
		_, exists := s.subscriptions[eventName]
		if !exists {
			s.subscriptions[eventName] = eventProcessorSlice{processor}
		} else {
			s.subscriptions[eventName] = append(s.subscriptions[eventName], processor)
		}
	}
}

func (s *simulation) Duration() int {
	return s.queue.Clock()
}

func (s *simulation) EventCount() int {
	return s.eventCount
}

func (s *simulation) routeEvent(event event.Event) {
	subscribers, exists := s.subscriptions[event.Name()]
	if !exists {
		log.Warningf("No subscribers for event: %s", event.Name())
		return
	}

	for _, processor := range subscribers {
		processor.Process(event, s.queue)
	}
}

func (s *simulation) Run(timeLimit int) {
	for {
		event := s.queue.Pop()
		s.routeEvent(event)

		s.eventCount++

		if s.shouldStop {
			break
		}
	}
}
