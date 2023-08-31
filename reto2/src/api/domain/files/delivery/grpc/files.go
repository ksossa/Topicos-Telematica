package grpc

import (
	"context"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/files/usecase"
	"log"
	"strings"
	"time"

	"github.com/ksossa/Topicos-Telematica/reto2/src/api/domain/proto/files"
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"
)

type FilesGRPCHandler struct {
	files.UnimplementedFilesServer
	usecase usecase.FileUseCases
}

func NewFilesGRPCHandler(container *dependencies.Container) *FilesGRPCHandler {
	return &FilesGRPCHandler{
		usecase: usecase.NewFileUseCases(container),
	}
}

func (handler FilesGRPCHandler) ReadAll(context.Context, *files.ReadAllRequest) (*files.ReadAllResponse, error) {
	log.Println(time.Now().String())
	_, err := handler.usecase.ReadAllMicro()
	if err != nil {
		return &files.ReadAllResponse{}, nil
	}
	mock := []string{"file1.txt, file2.txt"}
	ans := strings.Join(mock, ", ")
	log.Println(ans)
	return &files.ReadAllResponse{Files: ans}, nil
}
