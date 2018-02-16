package simulation

import (
	"testing"

	"github.bus.zalan.do/mboye/des/event"
	"github.com/stretchr/testify/assert"
)

type dummyEventProcessor struct {
}

func (p *dummyEventProcessor) Name() string {
	return "Dummy event processor"
}

func (p *dummyEventProcessor) Process(event event.Event, scheduler event.EventScheduler) {

}

func (p *dummyEventProcessor) Subscriptions() []string {
	return []string{"dummyEvent"}
}

func TestSimulationWithNoEvents(t *testing.T) {
	sim := NewSimulation()
	processor := &dummyEventProcessor{}

	sim.AddEventProcessor(processor)

	sim.Run(0)

	assert.Equal(t, 0, sim.Duration())
	assert.Equal(t, 1, sim.EventCount())
}
