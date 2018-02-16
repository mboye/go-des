package event

type scheduledEvent struct {
	scheduledAt int
	id          int
	event       Event
}

// Less indicates which of two events should be executed before the other
func Less(i, j scheduledEvent) bool {
	if i.scheduledAt < j.scheduledAt {
		return true
	} else if i.scheduledAt == j.scheduledAt {
		if i.id == j.id {
			panic("Event IDs must be unique")
		} else {
			return i.id < j.id
		}
	}

	return false
}
