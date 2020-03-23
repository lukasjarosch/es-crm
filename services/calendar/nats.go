package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/lukasjarosch/escrm/services/calendar/domain"
)

type NatsEventStore struct {
	nc                          *nats.Conn
	calendarCreatedSubscription *nats.Subscription
	calendarCreatedChan         chan domain.CalendarCreatedMessage
}

func NewNatsEventStore(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NatsEventStore{nc: nc}, nil
}

func (es *NatsEventStore) Close() {
	if es.nc != nil {
		es.nc.Close()
	}
	if es.calendarCreatedSubscription != nil {
		if err := es.calendarCreatedSubscription.Unsubscribe(); err != nil {
			log.Fatal(err)
		}
	}
	close(es.calendarCreatedChan)
}

func (es *NatsEventStore) PublishCalendarCreated(calendar domain.Calendar) error {
	evt := domain.CalendarCreatedMessage{
		ID:        calendar.ID,
		Name:      calendar.Name,
		CreatedAt: time.Time{},
	}
	data, err := es.writeEvent(evt)
	if err != nil {
	    return err
	}

	return es.nc.Publish(evt.Key(), data)
}

func (es *NatsEventStore) SubscribeCalendarCreated() (<-chan domain.CalendarCreatedMessage, error) {
	var err error
	es.calendarCreatedChan = make(chan domain.CalendarCreatedMessage)
	ch := make(chan *nats.Msg)
	evt := domain.CalendarCreatedMessage{}
	es.calendarCreatedSubscription, err = es.nc.ChanSubscribe(evt.Key(), ch)
	if err != nil {
	    return nil, err
	}

	// decode
	go func() {
		for {
			select {
			case msg := <-ch:
				if err := es.readEvent(msg.Data, &evt); err != nil {
					log.Fatal(err)
				}
				es.calendarCreatedChan <- evt
			}
		}
	}()
	return es.calendarCreatedChan, nil
}

func (es *NatsEventStore) OnCalendarCreated(f func(domain.CalendarCreatedMessage)) (err error) {
	evt := domain.CalendarCreatedMessage{}
	es.calendarCreatedSubscription, err = es.nc.Subscribe(evt.Key(), func(msg *nats.Msg) {
		log.Fatal(err)
	})
	f(evt)

	return
}

func (es *NatsEventStore) writeEvent(evt domain.Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(evt)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (es *NatsEventStore) readEvent(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}