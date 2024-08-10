package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cshep4/grpc-course/09-todo-service/proto"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewTodoServiceClient(conn)

	task1, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "wake up"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added task - id: %s", task1.GetId())

	task2, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "walk the dog"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added task - id: %s", task2.GetId())

	tasks, err := client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("existing tasks: %v", tasks.GetTasks())

	_, err = client.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: task1.GetId()})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("completed task - id: %s", task1.GetId())

	task3, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "have breakfast"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added task - id: %s", task3.GetId())

	tasks, err = client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("existing tasks: %v", tasks.GetTasks())
}