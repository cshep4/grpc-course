package stream

import (
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cshep4/grpc-course/module3-exercise/proto"
)

type Service struct {
	*proto.UnimplementedFileUploadServiceServer
}

func (s Service) DownloadFile(req *proto.DownloadFileRequest, stream proto.FileUploadService_DownloadFileServer) error {
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "filename cannot be empty")
	}

	// open file
	file, err := os.Open(req.Name)
	if err != nil {
		if os.IsNotExist(err) {
			return status.Error(codes.NotFound, "file not found")
		}
		return err
	}
	defer file.Close()

	const bufferSize = 5 * 1024 // send in 5KiB chunks
	buff := make([]byte, bufferSize)
	for {
		// read bytes
		bytesRead, err := file.Read(buff)
		if err != nil {
			if err == io.EOF {
				return nil // end of file, close stream
			}
			return err
		}

		// stream bytes to client
		err = stream.Send(&proto.DownloadFileResponse{Content: buff[:bytesRead]})
		if err != nil {
			return err
		}
	}
}

func (s Service) UploadFile(stream proto.FileUploadService_UploadFileServer) error {
	// generate file name
	fileName := fmt.Sprintf("%s.png", uuid.New().String())

	// create file
	fo, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fo.Close()

	for {
		// receive chunk
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// close stream and send response
				return stream.SendAndClose(&proto.UploadFileResponse{Name: fileName})
			}
			return err
		}

		// write chunk to file
		if _, err := fo.Write(res.Content); err != nil {
			return fmt.Errorf("failed to write to file %w", err)
		}
	}
}
