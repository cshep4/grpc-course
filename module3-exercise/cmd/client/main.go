package main

import (
	"log"
	"net/http"

	"github.com/cshep4/grpc-course/module3-exercise/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// initialise a gRPC connection on server start
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewFileUploadServiceClient(conn)

	http.HandleFunc("/", downloadHandler(client))

	log.Printf("starting http server on address: %s", ":8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// downloadHandler is an example of a gRPC client making a request to a server streaming RPC.
// The gRPC call will stream a file in chunks back to the client.
// The file content will be buffered until the server stream is complete, then the content will be returned to the user.
func downloadHandler(client proto.FileUploadServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// your implementation goes here ...
		panic("implement me")
	}
}
