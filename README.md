# Building Production-Ready Services with gRPC and Go

gRPC is the modern, high performance way to communicate across services. It is used by huge companies such as Google, Cloudflare and Netflix and, after this course, maybe by your company too!

Go is the perfect tool for gRPC, and by the end of this extensive course, you will see why.

This comprehensive course is designed for engineers to become proficient at writing production-ready microservices in Go using gRPC. Through a series of structured modules, students will learn to build scalable, robust and type-safe gRPC clients/services that are professional-grade. The course starts with the basics of Protocol Buffers and gRPC, gradually moving to advanced features, such as streaming, authentication/authorization and deployment. This code pairs with the [ByteSizeGo video course](https://www.bytesizego.com/grpc-with-go).

## Course Outline

### Module 1: Introduction to Protobuf, gRPC & protoc

Introduction into protobuf, use cases and benefits of using it. Discussing basic syntax and concepts, such as message & service definitions, running through data types, and language specific options.

- What are protocol buffers & basic concepts
- What is gRPC?
- Introduction to the Protoc tool & generating code

### Module 2: Building a Simple gRPC Service

Talk through service definitions in protobuf. Walk through creating a simple hello world application in Go and implement the client and server. Discuss errors and status codes and how to handle them in Go.

- Defining a gRPC service contract
- How to implement basic clients & servers
- How to handle errors in gRPC, including the different status codes.
- Practical example and exercise to build a simple application.

### Module 3: gRPC streaming

Discuss HTTP/2, multiplexing, plus benefits & use cases of streaming.

- The different types of streaming available in gRPC.
- How to implement server streaming.
- How to implement client streaming.
- How to implement bi-directional streaming.
- Exercise to build a streaming application.

### Module 4: Authentication, SSL/TLS

Authentication in gRPC. How SSL/TLS works and how it can be used with gRPC.

- What is SSL/TLS?
- Difference between server side TLS and mTLS.
- Implementing server side TLS in gRPC.
- Implementing mTLS in gRPC.
- Exercise to add mTLS authentication to an existing application.

### Module 5: Interceptors, Metadata & Authorisation

Discuss more advanced features that help enable things like authorisation & user session management.

- What are interceptors, getting RPC information in interceptor & chaining multiple.
- How to set deadlines/timeouts.
- CallOptions vs Metadata vs Context.
- Creating JWT authorization in a gRPC application.

### Module 6: Load Balancing in gRPC

Client-side load balancing in gRPC & other functionality provided by client service config.

- What is client service config, and how it can be used.
- Overview of client-side load balancing, load balancing strategies and implementation in gRPC.
- Configuring timeouts via client service config.
- Enabling automatic retry policies on RPC calls.

### Module 7: Testing gRPC services

The different types of tests in a gRPC service - both unit tests and integration tests, as well as tooling used to debug and test RPCs locally.

- Making gRPC requests locally (via Postman or grpccurl)
- Unit testing unary & streaming RPCs
- Creating integration/end-to-end tests for gRPC services

### Module 8: Containerising & Deploying to Kubernetes

Step through how to containerise a Go service in Docker, as well as and resources required to deploy to Kubernetes and expose to the internet.

- What is Docker & Kubernetes? How to build a Go service docker container.
- Creating a deployment & running on Kubernetes.
- How to securely expose gRPC services to the world.
- Automatic TLS certificate renewal with LetsEncrypt

### Module 9: Using Buf to manage protobuf

Discuss challenges of using protobuf as companies grow. Walk through tooling to manage protobuf effectively & introduce Buf.

- Challenges with maintaining a large number of protobufs.
- Building a protobuf registry as a centralised repository for your gRPC APIs.
- Ensure high standards through linting, formatting and breaking change detection.
- What is the ConnectRPC framework?
- How to manage protobuf dependencies using the Buf Schema Registry.
