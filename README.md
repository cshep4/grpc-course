# Modules

## Module 1: Introduction to Protobuf, gRPC & protoc

Introduction into protobuf, use cases and benefits of using it. Discussing basic syntax and concepts, such as message &
service definitions, running through data types, and language specific options.

Also introduce gRPC, discuss what it is and why it was created. Talk through use cases and benefits over rest. Talk
about different types of API and workflow for creating services.

Walk through protoc tool and how to download plugins to generate code in required languages.

- What are protocol buffers and language interoperability
- Basic protobuf concepts
    - Use cases and benefits
    - Message definitions
    - Primitive data types (string, int, float, bool, etc)
    - Commonly imported data types (timestamp, struct, empty)
    - Options (go_package)
    - proto2 vs proto3
- Protoc tool
    - Installing protoc
    - Go plugins for protoc
    - Generating code
    - Walking through generated code
- What is gRPC + gRPC workflow
- Why use gRPC?
    - Strong message structure
    - HTTP/2
    - Supports different languages
    - Faster message transmission
- gRPC vs REST
- Types of gRPC APIs
    - Unary
    - Server streaming
    - Client streaming
    - Bi-directional streaming

### Videos/Lessons:

1. What are protocol buffers and language interoperability
2. Basic protobuf concepts
3. What is gRPC + gRPC workflow
4. Why use gRPC? gRPC vs REST
5. Types of gRPC APIs
6. Protoc tool & generating code

### Exercise: Generate file protobuf message definitions in Go and import in a simple application

```bash
$ protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/hello.proto

$ go run main.go Chris
```

- Objective:
    - Install protoc and succesfully generate a basic message definition struct in Go. Create a simple hello world
      application that imports the protobuf code, instantiates an object and prints the result
- Requirements:
    - The protobuf file should include a message definition which contains a string field.
    - Should define a custom go package in the protobuf.
    - Should generate the go code into a subdirectory inside the pkg package.
    - Create a simple go script that accepts a name
    - Go script should import the generated protobuf code, instantiate an object using the inputted name and log the
      hello world message

## Module 2: Project overview & setup - building a simple gRPC service

Talk through service definitions in protobuf. Walk through creating a simple hello world application in Go and implement
the client and server. Discuss errors and status codes and how to handle them in Go.

- Defining a gRPC service in protobuf
- Go Dependencies Setup
- Implementing the server
- Custom error checking and status codes
- Registering the service and listening on a custom port
- Implementing the client
- Making gRPC calls
- Error handling

### Videos/Lessons:

1. Defining a gRPC service
2. Implementing a server
3. Implementing the client & making gRPC calls
4. Error Handling - client & server

### Exercise: Create a todo application

```protobuf
service TodoService {
  rpc AddTask(AddTaskRequest) returns (AddTaskResponse);
  rpc CompleteTask(CompleteTaskRequest) returns (CompleteTaskResponse);
  rpc ListTasks(ListTasksRequest) returns (ListTasksResponse);
}

message AddTaskRequest {
  string task = 1;
}

message AddTaskResponse {
  string id = 1;
}

message CompleteTaskRequest {
  string id = 1;
}

message CompleteTaskResponse {
}

message ListTasksRequest {
}

message ListTasksResponse {
  repeated Task tasks = 1;
}

message Task {
  string id = 1;
  string task = 1;
}
```

```bash
$ protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/todo.proto

$ go run cmd/server/main.go
$ go run cmd/client/main.go
```

- Objective:
    - Create a gRPC todo application. Implement both server and client using the same generated protobuf code and make
      gRPC calls.
- Requirements:
    - Protobuf should contain a single service defintion called TodoService.
    - Will contain the following RPCs:
        - AddTask
            - AddTaskRequest message should contain a task string
            - AddTaskResponse message should contain a generated task ID
            - Will add a task to the todo list
            - Server should return InvalidArgument if task is empty
        - CompleteTask
            - CompleteTaskRequest message contains a task ID
            - CompleteTaskResponse is an empty response
            - Will remove a task from the todo list
            - Server should return InvalidArgument if task ID is empty
            - Server should return NotFound if task ID is not found
        - ListTasks
            - ListTasksRequest message is an empty request
            - ListTasksResponse message contains a list of outstanding tasks
    - Server should listen on port 50051
    - Client should initialise gRPC connection, add tasks to todo list, list tasks and remove tasks and log results

## Module 3: gRPC streaming

Discuss HTTP/2, multiplexing and benefits & use cases of streaming.

- What is HTTP/2, multiplexing and gRPC streaming?
- When streaming should be used (and when it shouldn’t)
- Defining streaming RPCs in protobuf
- Implementing server streaming - server
- Implementing server streaming - client
- Implementing client streaming - server
- Implementing client streaming - client
- Implementing bi-directional streaming - server
- Implementing bi-directional streaming - client

### Videos/Lessons:

1. What is HTTP/2, multiplexing and gRPC streaming? - (word of caution when streaming should/shouldn't be used)
2. Implementing server streaming - client/server
3. Implementing client streaming - client/server
4. Implementing bi-directional streaming - client/server

### Exercise: Creating a file upload application

```protobuf
service FileUploadService {
  rpc DownloadFile(DownloadFileRequest) returns (stream DownloadFileResponse);
  rpc UploadFile(stream UploadFileRequest) returns (stream UploadFileResponse);
}

message DownloadFileRequest {
  string name = 1;
}

message DownloadFileResponse {
  bytes content = 1;
}

message UploadFileRequest {
  bytes content = 1;
  bool has_more = 2;
}

message UploadFileResponse {
  string name = 1;
}
```

- Objective:
    - Implement a basic file upload service that contains two RPCs:
- UploadFile - A bi-directional streaming RPC, that uploads a file in chunks to the server, whilst the server streams
  back updates when the chunk is successfully stored.
- DownloadFile - A server streaming RPC that accepts a filename in the request and streams the file content in chunks
  back to the client.
- Requirements:
    - Both RPCs should be implemented
    - When the server receives a message, it should push to clients subscribed to the chat room, with a different
      username
    - Messages should not be pushed to the client that sent the message (e.g. clients with the same username)
    - Messages do not need to be cached/persisted
    - When a client unsubscribes, the username should be removed from the list of recipients when messages are sent

## [BONUS] Module 3.1: Using gRPC Web

Talking through gRPC web, how this can be used in a browser, walk through code generation and setup, with simple example

- What is gRPC web?
- Required dependencies for using in a browser
- Generating gRPC web stubs
- Implementing a file upload client

### Videos/Lessons:

1. What is gRPC web?
2. Generating gRPC web stubs & required dependencies
3. Implementing a client

### Exercise: gRPC Implementation for file upload application

- Objective:
    - Implement gRPC web client logic for provided UI for file upload application
- Requirements:
    - Simple UI will be provided to upload and download files from server, with unimplemented API layer
    - Initialise gRPC connection & client stub for file upload service
    - Implement method to upload files in chunks to server using bi-directional streaming RPC
    - Implement method to download file in chunks given a file name by making calls to DownloadFile RPC

## Module 4: Authentication, SSL/TLS

Discuss authentication in gRPC. How SSL/TLS works and how it can be used with gRPC.

- What is SSL/TLS?
- Difference between server side TLS and mTLS
- Implementing server side TLS in gRPC
- Implementing mTLS in gRPC
- Automatic cert renewal using LetsEncrypt

### Videos/Lessons:

1. What is SSL/TLS?
2. What is mTLS?
3. Implementing server side TLS in gRPC
4. Implementing mTLS in gRPC
5. Automatic cert renewal using LetsEncrypt

### Exercise: Add mTLS to file upload application

- Objective:
    - Add mTLS authentication on both client & server side of the file sharing app
- Requirements:
    - Generate TLS certs for client & server
    - Use certs when initialising gRPC connection
    - Should use mTLS on each gRPC call

## Module 5: Interceptors, Metadata & Authorisation

Discuss more advanced features that help enable things like authorisation & user session management.

- Interceptors - introduction, client interceptors, server interceptors
- Getting RPC Information In Interceptor
- Chaining multiple interceptors
- Deadlines/timeouts
- CallOptions vs Metadata vs Context
- API key validator via interceptors
- Propagating user info via context
- Auth via CallCredentials

### Videos/Lessons:

1. Interceptors - introduction, client interceptors, server interceptors
2. Setting Deadlines/timeouts
3. CallOptions vs Metadata vs Context
4. Authorisation via interceptors & propagating user info via context
5. Auth via CallCredentials

### Exercise: Add interceptor to parse JWT information & propagate via context

```protobuf

service TokenService {
  rpc Validate(ValidateRequest) returns (ValidateResponse);
}

message ValidateRequest {
}

message ValidateResponse {
  map<string, string> claims = 1;
}
```

- Objective:
    - Implement gRPC service with interceptor to read token from metadata, parse the claims & propagate via context
- Requirements:
    - Implement provided protobuf service
    - Service should have custom interceptor
    - Gets token from middleware, if none is provided it should return PermissionDenied
    - Parse the claims from the token and add to context
    - In the RPC implementation, read the claims from the context and return in response

## Module 6: Load Balancing in gRPC

Short module to discuss client-side load balancing in gRPC. This can be used to distribute load across multiple server
instances. Will discuss how this can be used alongside tools such as Consul.

- Overview of client-side load balancing
- Service discovery
- Load balancing strategies
- Resilience and health checks
- Performance implications/considerations
- Example of load balancing in gRPC:
    - [https://pastebin.com/B2G9VPvv](https://pastebin.com/B2G9VPvv)

### Videos/Lessons:

1. Overview of client-side load balancing
2. Health checking & service discovery
3. Load balancing strategies
4. Performance implications/considerations
5. Example of load balancing in gRPC client

## Module 7: Testing gRPC services

Discuss the different types of tests in a gRPC service - both unit tests and integration tests, as well as tooling used
to debug and test RPCs locally

- Making gRPC requests locally (via Postman or grpccurl)
- Unit testing unary RPCs
- Unit testing streaming RPCs
- Creating integration/end-to-end tests for gRPC services

### Videos/Lessons:

1. Making gRPC requests via Postman
2. Making gRPC requests via grpccurl
3. Creating unit tests for RPCs
4. Creating end-to-end tests

### Exercise: Add tests to an existing application

- Objective:
    - Implement provided test suites for unit & end-to-end tests for an existing application
- Requirements:
    - 2 unimplemented test suites will be provided, for unit tests & end-to-end tests, along with a sample application
    - Implement the tests so that they sufficiently cover the logic & all pass

## Module 8: Containerising & Deploying to Kubernetes

Step through how to containerise a Go service in Docker, along with the steps and resources required to deploy to
Kubernetes and expose to the internet

- What is Docker & Kubernetes?
- A sample dockerfile to build a Go service
- Pushing docker image to a container registry
- Required resources to deploy a service in K8s
- Creating a Deployment
- Creating a Service to make the application discoverable within K8s
- Creating an Ingress & how to expose publicly
- TLS & cert issuer using LetsEncrypt

### Videos/Lessons:

1. What is Docker & Building a Go service in Docker
2. What is Kubernetes & Required resources to deploy a service in K8s - deployment, service, ingress
3. TLS & cert issuer using LetsEncrypt

### Exercise: Dockerise application & deploy to a local Kubernetes cluster

- Objective:
    - Take an existing application, containerise it & push to a local k8s cluster
- Requirements:
    - Sample application will be provided
    - Create a multi-stage dockerfile to build the application & run it
    - Create required resources to deploy to K8s (can use local cluster within Docker Desktop)
    - Create ingress to expose outside of cluster & validate by making requests from Postman

## Module 9: Using Buf to manage protobuf

Discuss challenges of using protobuf as companies grow. Talk about Buf and its features, highlighting its ease of use
and reasons for using over vanilla protoc

- Challenges with maintaining a large number of protobufs
- Improving discoverability of services
- Improving quality through linting and formatting
- Issues with breaking changes and how to avoid them
- Buf schema registry as a package manager

### Videos/Lessons:

1. Challenges with maintaining a large number of protobufs & What is Buf?
2. Problem #1 - Discoverability of services
3. Problem #2 - Inconsistent quality
4. Problem #3 - Breaking changes
5. Problem #4 - Generating code in all languages
6. Example Protobuf Registry

### Exercise: Use Buf plugins to generate protobuf in Go & web

- Objective:
    - Add Buf workspace to protobuf directory and generate go code using Buf plugins and command
- Requirements:
    - Add Buf workspace to protobuf directory
    - Using go and gRPC go plugins to generate code
    - Generate using Buf generate command
    - Run fmt and lint commands

## Module 10: Connect RPC Framework

Introduction into the Connect RPC framework written by Buf. A simpler and more readable framework which creates
fully-compatible gRPC APIs. As it also supports both JSON and binary-encoded Protobuf, it's as easy to call as using
curl.

- What is Connect RPC?
- Benefits of Connect?
- Generating & walking through Connect code
- Example Connect server implementation
- Example Connect client implementation
- Using JSON encoded protobuf and calling via curl

### Videos/Lessons:

1. What is Connect RPC?
    1. Benefits of Connect?
    2. Generating & walking through Connect code
2. Example Connect server implementation
3. Example Connect client implementation
4. Using JSON encoded protobuf and calling via curl

### Exercise: Convert to do application to use Connect

- Objective:
    - Convert the to do application from Module 2 to use Connect
- Requirements:
    - Generate Connect code using Buf
    - Implement server
    - Implement client
    - Validate by making requests via curl