# Module 07 Exercise: Adding unit tests to our todo application

## Objective

Take the gRPC todo application we built in Module 2 and implement the provided test suites for unit tests.

A mock for the TodoStore interface has been provided [here](internal%2Fmocks%2Fstore%2Fmock_store.gen.go).

A makefile has been provided with useful commands.

## Requirements

- Implement the provided unit test suite in [service_test.go](internal%2Ftodo%2Fservice_test.go).
- The tests should sufficiently cover the logic within the [service.go](internal%2Ftodo%2Fservice.go).
- All tests should pass.

## Useful Commands

To install protoc and required plugins:

```bash
$ make get-protoc-plugins
```

To compile protocol buffers:

```bash
$ make proto-gen
```

To run server:

```bash
$ make run-server
```

To run client:

```bash
$ make run-client
```

To run tests:

```bash
$ make test-unit
```