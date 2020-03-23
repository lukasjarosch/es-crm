package domain

import "time"

type Message interface {
	Key() string
}

type CalendarCreatedMessage struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

func (c CalendarCreatedMessage) Key() string {
	return "calendar.created"
}
