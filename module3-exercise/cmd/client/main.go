package main

import (
	"io"
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
		ctx := r.Context()

		// make request to server and initialise stream
		stream, err := client.DownloadFile(ctx, &proto.DownloadFileRequest{
			Name: "gopher.png",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// create a slice to store file contents
		var fileContent []byte

		// read chunks from server and store in slice
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			log.Println("chunk received from server")

			fileContent = append(fileContent, res.GetContent()...)
		}

		log.Println("server stream done")

		// write our slice of bytes back to the client
		if _, err := w.Write(fileContent); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
