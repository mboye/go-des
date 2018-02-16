package event

// Event represents an event in a simulation
type Event interface {
	Execute()
	ScheduledAt() uint64
	ID() uint64
}

type event struct {
	scheduledAt uint64
	id          uint64
}

// Less indicates which of two events should be executed before the other
func Less(i, j Event) bool {
	if i.ScheduledAt() < j.ScheduledAt() {
		return true
	} else if i.ScheduledAt() == j.ScheduledAt() {
		if i.ID() == j.ID() {
			panic("Event IDs must be unique")
		} else {
			return i.ID() < j.ID()
		}
	}

	return false
}
