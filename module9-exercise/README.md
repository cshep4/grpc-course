# Module 09 Exercise: Create a todo application

## Objective

Take the gRPC todo application we built in Module 2 and modify it to use the Connect framework. A gRPC implementation of the todo service has been provided, with protobuf code generated through Buf.

You should modify the [buf.gen.yaml](buf.gen.yaml) file to generate the [Connect bindings](https://buf.build/connectrpc/go), then modify the server and client to use Connect.

A makefile has been provided with useful commands.

To install Buf:
```bash
$ make get-buf
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

## Requirements

- Generate Connect bindings using Buf.
- Connect-go plugins should be downloaded from the Buf Schema Registry - https://buf.build/connectrpc/go
- Modify the server to use the Connect framework.
- Modify the client to use the Connect framework.