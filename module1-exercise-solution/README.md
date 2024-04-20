# Module 01 Exercise: Generate file protobuf message definitions in Go

## Objective

Install protoc and successfully generate a basic message definition struct in Go. Create a simple hello world
application that takes a name as the first argument, imports the protobuf code, instantiates an object and prints the
result.

To install protoc and required plugins:

```bash
$ brew install protobuf
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33
```

To run the application:

```bash
$ go run main.go <name>
```

To compile protocol buffers:

```bash
$ protoc --go_out=. --go_opt=paths=source_relative proto/hello.proto
```

## Requirements

- The protobuf file should include a message definition which contains a string field.
- Define a custom go package in the protobuf.
- Generate the go code into a subdirectory inside the `proto` package.
- Create a simple go script that accepts a name
- Go script should import the generated protobuf code, instantiate an object using the inputted name and log the hello
  world message

## Example

Example input:

```bash
$ go run main.go Chris
```

Expected output:

```bash
2024/04/20 20:28:14 Hello Chris!
```
