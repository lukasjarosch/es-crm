module github.com/lukasjarosch/escrm

go 1.14

replace github.com/lukasjarosch/escrm/proto/gen/go => ../../proto/gen/go

require (
	github.com/cockroachdb/cockroach-go v0.0.0-20200312223839-f565e4789405 // indirect
	github.com/fatih/color v1.9.0
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/jackc/pgconn v1.4.0 // indirect
	github.com/magefile/mage v1.9.0
	github.com/nats-io/gnatsd v1.4.1 // indirect
	github.com/nats-io/go-nats v1.7.2 // indirect
	github.com/nats-io/go-nats-streaming v0.4.4 // indirect
	github.com/nats-io/nats-server v1.4.1 // indirect
	github.com/nats-io/nats-streaming-server v0.17.0 // indirect
	github.com/nats-io/nats.go v1.9.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shijuvar/go-distributed-sys v0.0.0-20180815023924-fe08be5bef7a // indirect
	google.golang.org/grpc v1.28.0 // indirect
)
