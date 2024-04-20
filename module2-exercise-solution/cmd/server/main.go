package main

import (
	"github.com/cshep4/grpc-course/02-todo-service/internal/todo"
	"github.com/cshep4/grpc-course/02-todo-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	todoService := todo.NewService()

	proto.RegisterTodoServiceServer(grpcServer, todoService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	log.Printf("starting grpc server on address: %s", ":50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
