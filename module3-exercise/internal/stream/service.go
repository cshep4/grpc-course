package stream

import (
	"fmt"
	"github.com/cshep4/grpc-course/module3-exercise/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
)

type Service struct {
	*proto.UnimplementedFileUploadServiceServer
}

func (s Service) DownloadFile(req *proto.DownloadFileRequest, stream proto.FileUploadService_DownloadFileServer) error {
	// check if the name is empty in the request
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "file name cannot be empty")
	}

	// open file
	file, err := os.Open(req.Name)
	if err != nil {
		// check if the is found - return not found otherwise
		if os.IsNotExist(err) {
			return status.Error(codes.NotFound, "file not found")
		}
		return err
	}

	// read the file in chunks of 5KB
	const bufferSize = 5 * 1024
	buff := make([]byte, bufferSize)
	for {
		bytes, err := file.Read(buff)
		if err != nil {
			if err == io.EOF {
				return nil // close the server stream
			}
			return status.Error(codes.Internal, "error reading file")
		}

		// stream each chunk to the client
		err = stream.Send(&proto.DownloadFileResponse{Content: buff[:bytes]})
		if err != nil {
			return err
		}
	}
}

func (s Service) UploadFile(stream proto.FileUploadService_UploadFileServer) error {
	// generate the file name
	fileName := fmt.Sprintf("%s.png", uuid.New().String())

	// create a file
	file, err := os.Create(fileName)
	if err != nil {
		return status.Error(codes.Internal, "error creating file")
	}
	defer file.Close()

	// receive chunks from client
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&proto.UploadFileResponse{
					Name: fileName,
				})
			}
			return err
		}

		// write chunk to file
		if _, err := file.Write(res.Content); err != nil {
			return status.Error(codes.Internal, "error writing to file")
		}
	}
}
