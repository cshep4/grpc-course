# Module 05 Exercise: Add interceptor to parse JWT information & propagate via context

## Objective

Implement a gRPC service with a custom interceptor to read token from metadata, parse the claims & propagate via
context. A [service contract](./proto/token.proto) has been provided.

A JWT has been provided in the [client](cmd/client/main.go) on line 15, with the following claims:

```json
{
  "exp": 4113023550,
  "iat": 1714735950,
  "name": "Chris",
  "role": "admin",
  "sub": "user-id-1234"
}
```

This should be passed as a header to a gRPC call to the `Validate` RPC. The server should have a custom interceptor
defined that does the following:

1. Gets the token from the incoming metadata.
2. Validates the token using the provided `JWT_SECRET`.
    1. You can use [golang-jwt/jwt](https://github.com/golang-jwt/jwt/) library for this.
3. Parses the claims and adds them to the context before calling the RPC handler.

The `Validate` RPC should then retrieve the claims map from the context, and return in the response.

The response from the RPC should look like this:

```json
{
  "exp": "4113023550",
  "iat": "1714735950",
  "name": "Chris",
  "role": "admin",
  "sub": "user-id-1234"
}
```

## Requirements:

- Implement provided protobuf service
- Service should have custom interceptor
- Should return `PermissionDenied` if no token exists in metadata
- Should return `PermissionDenied` if token is not valid
- Claims from the token should be parsed to a `map[string]string`
- Claims should be propagated to main RPC handler via context.
- In the RPC implementation, read the claims from the context and return in response

## Tip

You can use the debugger at [jwt.io](https://jwt.io) to check the token and view claims.

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