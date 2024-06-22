.PHONY: all
all: get-buf proto-gen

.PHONY: get-buf
get-buf:
	brew install bufbuild/buf/buf

.PHONY: proto-gen
proto-gen:
	buf generate proto

.PHONY: run-server
run-server:
	go run cmd/server/main.go

.PHONY: run-client
run-client:
	go run cmd/client/main.go