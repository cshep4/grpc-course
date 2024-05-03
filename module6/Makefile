.PHONY: all
all: get-protoc-plugins proto-gen

.PHONY: get-protoc-plugins
get-protoc-plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/config.proto

.PHONY: run-servers
run-servers:
	PORT=50051 go run cmd/server/main.go &
	PORT=50052 go run cmd/server/main.go &
	PORT=50053 go run cmd/server/main.go
