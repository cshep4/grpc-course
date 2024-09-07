package main

import (
	"context"
	"encoding/json"
	"github.com/cshep4/grpc-course/module6/internal/config"
	"github.com/cshep4/grpc-course/module6/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	ctx := context.Background()

	//	const grpcServiceConfig = `{
	//	"methodConfig": [{
	//		"name": [{"service": "config.ConfigService", "method": "LongRunning"}],
	//		"timeout": "1s"
	//	}]
	//}`
	cfg := config.Config{
		MethodConfig: []config.MethodConfig{{
			Name: []config.NameConfig{{
				Service: "config.ConfigService",
				Method:  "LongRunning",
			}},
			Timeout: "1s",
		}},
	}

	serviceConfig, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalf("failed to marshal config: %v", err)
	}

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(string(serviceConfig)),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewConfigServiceClient(conn)

	_, err = client.LongRunning(ctx, &proto.LongRunningRequest{})
	if err != nil {
		log.Fatal(err)
	}
}
