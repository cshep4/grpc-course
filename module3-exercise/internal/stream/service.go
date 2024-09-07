package stream

import (
	"github.com/cshep4/grpc-course/module3-exercise/proto"
)

type Service struct {
	*proto.UnimplementedFileUploadServiceServer
}

func (s Service) DownloadFile(req *proto.DownloadFileRequest, stream proto.FileUploadService_DownloadFileServer) error {
	// your implementation goes here ...
	panic("implement me")
}

func (s Service) UploadFile(stream proto.FileUploadService_UploadFileServer) error {
	// your implementation goes here ...
	panic("implement me")
}
