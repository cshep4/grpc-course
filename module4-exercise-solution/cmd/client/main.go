package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/cshep4/grpc-course/module4-exercise/proto"
)

func main() {
	// Load the client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		log.Fatalf("failed to load client certificate and key: %v", err)
	}

	// Load the CA's certificate to verify the server
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatalf("failed to load CA certificate: %v", err)
	}

	// append the CA's certificate to the cert pool
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Fatalf("failed to append CA certificate to pool")
	}

	// Create the TLS config for the client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool,
	}

	creds := credentials.NewTLS(tlsConfig)

	// initialise a gRPC connection
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(creds),
	)
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

		// make request to gRPC server and initialise server stream
		stream, err := client.DownloadFile(ctx, &proto.DownloadFileRequest{Name: "gopher.png"})
		if err != nil {
			// check status code returned from server
			st := status.Convert(err)
			switch st.Code() {
			case codes.NotFound:
				http.Error(w, "File not found.", 404)
				return
			case codes.InvalidArgument:
				http.Error(w, "Bad request.", 400)
				return
			}

			http.Error(w, err.Error(), 500)
			return
		}

		log.Println("server stream started")

		// create slice of file contents
		var fileContents []byte

		for {
			// receive file chunk
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break // stream done
				}

				http.Error(w, err.Error(), 500)
				return
			}

			log.Println("received file chunk")

			// append file chunk to slice
			fileContents = append(fileContents, res.Content...)
		}

		log.Println("server stream done")

		// return file contents to user
		if _, err := w.Write(fileContents); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}
