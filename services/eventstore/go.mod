module eventstore

go 1.14

replace github.com/lukasjarosch/escrm/protobuf-go => ../../proto/gen/go

require (
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/lib/pq v1.3.0
	github.com/lukasjarosch/escrm/protobuf-go v0.0.0-00010101000000-000000000000
	github.com/nats-io/go-nats v1.7.2 // indirect
	github.com/nats-io/go-nats-streaming v0.4.4
	github.com/nats-io/nkeys v0.1.3 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/shijuvar/go-distributed-sys v0.0.0-20180815023924-fe08be5bef7a
	google.golang.org/grpc v1.28.0
)
