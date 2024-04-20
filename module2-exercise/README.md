# Module 02 Exercise: Create a todo application

## Objective

Create a gRPC todo application. A [service contract](./proto/todo.proto) has been provided. You should compile the
protobuf to generate the client & server stubs, then implement both server and client using the same generated protobuf
code and make gRPC calls.

## Requirements

- Implement all the defined RPCs on the [TodoService contract](./proto/todo.proto).
- The server should store todos in-memory in a map.
- Each RPC should:
    - AddTask
        - Will add a task to the todo map
        - AddTaskRequest message contains a task string
        - AddTaskResponse message contains a generated task ID
        - Server should return InvalidArgument if task is empty
    - CompleteTask
        - Will remove a task from the todo map
        - CompleteTaskRequest message contains a task ID
        - CompleteTaskResponse is an empty response
        - Server should return InvalidArgument if task ID is empty
        - Server should return NotFound if task ID is not found
    - ListTasks
        - Will return a list of outstanding tasks
        - ListTasksRequest message is an empty request
        - ListTasksResponse message contains a list of outstanding tasks
- Server should listen on port 50051
- Client should initialise gRPC connection, add tasks to todo list, list tasks and remove tasks and log results

## Commands

A makefile has been provided with useful commands.

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