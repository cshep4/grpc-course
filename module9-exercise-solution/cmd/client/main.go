package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"

	"github.com/cshep4/grpc-course/09-todo-service/proto"
	"github.com/cshep4/grpc-course/09-todo-service/proto/protoconnect"
)

func main() {
	ctx := context.Background()

	client := protoconnect.NewTodoServiceClient(http.DefaultClient, "http://localhost:50051")

	task1, err := client.AddTask(ctx, connect.NewRequest(&proto.AddTaskRequest{Task: "wake up"}))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added task - id: %s", task1.Msg.GetId())

	task2, err := client.AddTask(ctx, connect.NewRequest(&proto.AddTaskRequest{Task: "walk the dog"}))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added task - id: %s", task2.Msg.GetId())

	tasks, err := client.ListTasks(ctx, connect.NewRequest(&proto.ListTasksRequest{}))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("existing tasks: %v", tasks.Msg.GetTasks())

	_, err = client.CompleteTask(ctx, connect.NewRequest(&proto.CompleteTaskRequest{Id: task1.Msg.GetId()}))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("completed task - id: %s", task1.Msg.GetId())

	task3, err := client.AddTask(ctx, connect.NewRequest(&proto.AddTaskRequest{Task: "have breakfast"}))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added task - id: %s", task3.Msg.GetId())

	tasks, err = client.ListTasks(ctx, connect.NewRequest(&proto.ListTasksRequest{}))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("existing tasks: %v", tasks.Msg.GetTasks())
}
