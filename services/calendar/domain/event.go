package domain

type EventStore interface {
	Close()
	PublishCalendarCreated(calendar Calendar) error
	SubscribeCalendarCreated() (<-chan CalendarCreatedMessage, error)
	OnCalendarCreated(f func(CalendarCreatedMessage)) error
}

var impl EventStore

func SetEventStore(es EventStore) {
	impl = es
}

func Close() {
	impl.Close()
}

func PublishCalendarCreated(created Calendar) error {
	return impl.PublishCalendarCreated(created)
}

func SubscribeCalendarCreated() (<-chan CalendarCreatedMessage, error) {
	return impl.SubscribeCalendarCreated()
}

func OnCalendarCreated(f func(CalendarCreatedMessage)) error {
	return impl.OnCalendarCreated(f)
}
