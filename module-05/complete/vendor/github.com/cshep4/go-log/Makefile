.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

default: mod test

LINTER_VERSION=v1.23.3

get-linter:
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin ${LINTER_VERSION}

lint: get-linter
	golangci-lint run --timeout=5m

get-generator:
	go install github.com/golang/mock/mockgen

generate: get-generator
	go generate -mod vendor ./...

clean-mock:
	find internal -iname '*_mock.go' -exec rm {} \;

regenerate: clean-mock generate

mod:
	go mod vendor -v

test: lint
	go test -v -mod vendor -race -coverprofile=race.out ./...

cover-ci:
	go tool cover -func=race.out

cover:
	go test -race ./... -vet all -coverprofile=coverage.out
	go tool cover -html=coverage.out


