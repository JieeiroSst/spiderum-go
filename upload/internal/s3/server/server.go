package server

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/upload/config"
	"gitlab.com/Spide_IT/spide_it/upload/internal/s3/delivery"
	"gitlab.com/Spide_IT/spide_it/upload/internal/s3/repository"
	"gitlab.com/Spide_IT/spide_it/upload/internal/s3/usecase"
)

type UploadServer struct {
	server *gin.Engine
	s3 *session.Session
	cgf *config.Config
}

func NewUploadServer(server *gin.Engine, s3 *session.Session) *UploadServer {
	return &UploadServer{
		server:server,
		s3:s3,
	}
}

func (uploadServer *UploadServer) Run(){
	repo:=repository.NewRepository(uploadServer.cgf,uploadServer.s3)
	use :=usecase.NewUsecase(repo)
	delivery.NewHttp(uploadServer.server,use,uploadServer.s3)
}