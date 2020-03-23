package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	eventstorev1 "github.com/lukasjarosch/escrm/protobuf-go/crm/eventstore/v1"
	_ "github.com/lib/pq"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/shijuvar/go-distributed-sys/natsutil"
	"google.golang.org/grpc"
)

const (
	port      = ":50051"
	clusterID = "test-cluster"
	clientID  = "event-store"
)

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("postgres", "postgresql://eventstore@localhost:26257/eventstore?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS event_store (
    			id string PRIMARY KEY, 
    			event_type string, 
    			aggregate_id string, 
    			aggregate_type string, 
    			event_payload string, 
    			timestamp_us bigint,
    			channel string)`)
	if err != nil {
	    log.Fatal(err)
	}
}

type grpcServer struct {
	*natsutil.StreamingComponent
}

func (s *grpcServer) CreateEvent(ctx context.Context, req *eventstorev1.CreateEventRequest) (*eventstorev1.CreateEventResponse, error) {
	es := EventStore{}
	err := es.CreateEvent(req.Event)
	if err != nil {
	    return nil, err
	}

	cs := s.StreamingComponent.NATS()
	channel := req.Event.Channel
	eventMsg := []byte(req.Event.Payload)
	err = cs.Publish(channel, eventMsg)
	if err != nil {
	    log.Fatal(err)
	}
	log.Printf("published event on '%s'\n", channel)

	return &eventstorev1.CreateEventResponse{}, nil
}

func (s *grpcServer) GetEvents(ctx context.Context, req *eventstorev1.GetEventsRequest) (*eventstorev1.GetEventsResponse, error) {
	return &eventstorev1.GetEventsResponse{}, nil
}

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
	    log.Fatal(err)
	}

	comp := natsutil.NewStreamingComponent(clientID)

	err = comp.ConnectToNATSStreaming(clusterID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
	    log.Fatal(err)
	}

	srv := grpc.NewServer()
	eventstorev1.RegisterEventStoreAPIServer(srv, &grpcServer{StreamingComponent: comp})
	srv.Serve(list)
}