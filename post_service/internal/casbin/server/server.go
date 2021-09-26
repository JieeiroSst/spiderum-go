package server

import (
	"github.com/gin-gonic/gin"
	delivery "gitlab.com/Spide_IT/spide_it/internal/casbin/delivery/http"
	"gitlab.com/Spide_IT/spide_it/internal/casbin/http"
	"gitlab.com/Spide_IT/spide_it/internal/casbin/repository"
	"gitlab.com/Spide_IT/spide_it/internal/casbin/usecase"
	"gorm.io/gorm"
)

type CasbinServer struct {
	server *gin.Engine
	db *gorm.DB
}

func NewCasbinServer(server *gin.Engine,db *gorm.DB) *CasbinServer{
	return &CasbinServer{
		server:server,
		db:db,
	}
}

func (casbinServer *CasbinServer) Run(){
	repo:=repository.NewCasbinRuleRepository(casbinServer.db)
	useCase :=usecase.NewCasbinRuleUseCase(repo)
	http:=http.NewHttp(useCase)
	delivery.NewHttp(casbinServer.server,*http)
}