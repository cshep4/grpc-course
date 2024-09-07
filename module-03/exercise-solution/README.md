# Module 03 Exercise: Create a file upload application

## Objective

Implement a basic file upload service for the provided protobuf contract, that contains two RPCs:

- UploadFile - A client streaming RPC, that uploads a file in chunks to the server. Once uploaded, the server responds
  with a generated file name.
- DownloadFile - A server streaming RPC that accepts a filename in the request and streams the file content in chunks
  of 5KB back to the client.

Once the server is created, implement the client which will call the `DownloadFile` RPC to download the
provided [gopher.png](./gopher.png) image. A basic HTTP server has been provided to test the client inside a browser,
you just need to implement the `downloadHandler` in the [main.go](./cmd/client/main.go) file to download the file and
return the content to the user.

Once the client application is running, you should be able to test it in a browser by going to: http://localhost:8080

## Requirements:

- Both RPCs should be implemented on the server.
- Server should listen on port 50051.
- Server should return `InvalidArgument` in the DownloadFile RPC if the file name is empty.
- Server should return `NotFound` in the DownloadFile RPC if the file is not found.
- Server should download the file in chunks of 5KB.
- Implement `downloadHandler` function in client to download the [gopher.png](./gopher.png) file in chunks using your
  gRPC service.
- Client HTTP server should buffer the file in memory until all chunks are received, then return the file response.

## Useful Commands

A makefile has been provided with useful commands.

To install protoc and required plugins:

```bash
$ make get-protoc-plugins
```

To compile protocol buffers:

```bash
$ make proto-gen
```

To run server application:

```bash
$ make run-server
```

To run client application:

```bash
$ make run-client
```