package main

import (
	eventstorev1 "github.com/lukasjarosch/escrm/protobuf-go/crm/eventstore/v1"
	"github.com/pkg/errors"
)

type EventStore struct{}

func (s EventStore) CreateEvent(event *eventstorev1.Event) error {
	sql := `INSERT INTO event_store (id, event_type, aggregate_id, aggregate_type, event_payload, timestamp_us, channel) VALUES ($1, $2, $3, $4, $5, $6, $7);`
	_, err := db.Exec(sql, event.Id, event.Type, event.AggregateId, event.AggregateType, event.Payload, event.TimestampUs, event.Channel)
	if err != nil {
		return errors.Wrap(err, "CreateEvent")
	}
	return nil
}

func (s EventStore) GetEvents(filter *eventstorev1.EventFilter) error {
	return nil
}