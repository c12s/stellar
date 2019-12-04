module github.com/c12s/stellar

replace github.com/coreos/go-systemd/journal => ../../coreos/go-systemd/journal

go 1.13

require (
	github.com/c12s/scheme v0.0.0-20191204214602-7126694c68c7
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/go-systemd/journal v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1 // indirect
	github.com/nats-io/go-nats v1.7.3-0.20190608183121-73ffc26dfe70
	github.com/nats-io/nats.go v1.9.1 // indirect
	github.com/nats-io/nkeys v0.1.3 // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20191202143827-86a70503ff7e // indirect
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	google.golang.org/grpc v1.25.1
	gopkg.in/yaml.v2 v2.2.7
)
