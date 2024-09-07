.PHONY: all
all: get-buf proto-gen

.PHONY: get-buf
get-buf:
	brew install bufbuild/buf/buf

.PHONY: mod
mod:
	buf dep update proto

.PHONY: proto-gen
proto-gen:
	buf generate proto

.PHONY: run-grpc-server
run-grpc-server:
	go run cmd/grpc-server/main.go

.PHONY: run-grpc-client
run-grpc-client:
	go run cmd/grpc-client/main.go

create-client:
	npx create-next-app client
