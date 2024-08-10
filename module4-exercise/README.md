# Module 04 Exercise: Add mTLS to our file upload application

## Objective

Take the gRPC file upload service we built in module 3 and implement mTLS authentication.

You should generate the client and server certificates using the `gen-certs` make command provided. Once these are
generated, modify both the client and server to load these on startup, and ensure the certificates are verified when
making a connection.

Once the client application is running, you should be able to test it in a browser by going to: http://localhost:8080

If the image of a gopher is returned successfully then mTLS authentication is working as expected and communication is
secure.

## Requirements:

- Generate TLS certifcates for both client & server using the `gen-certs` make command.
- Certificates should be loaded and passed as credentials on server startup.
- Certificates should be loaded and used in TLS config when initialising gRPC connection on the client.

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

To generate all TLS certificates:

```bash
$ make gen-certs
```

To generate only the root CA certificates:

```bash
$ make gen-ca-certs
```

To generate only the server certificates:

```bash
$ make gen-server-certs
```

To generate only the client certificates:

```bash
$ make gen-client-certs
```