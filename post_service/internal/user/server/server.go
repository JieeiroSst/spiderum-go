package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/config"
	delivery "gitlab.com/Spide_IT/spide_it/internal/user/delivery/http"
	"gitlab.com/Spide_IT/spide_it/internal/user/http"
	"gitlab.com/Spide_IT/spide_it/internal/user/repository"
	"gitlab.com/Spide_IT/spide_it/internal/user/usecase"
	"gitlab.com/Spide_IT/spide_it/pkg/bigcache"
	"gitlab.com/Spide_IT/spide_it/pkg/jwt"
	"gitlab.com/Spide_IT/spide_it/pkg/snowflake"
	"gitlab.com/Spide_IT/spide_it/utils"
	"gorm.io/gorm"
)

type UserServer struct {
	server *gin.Engine
	config *config.Config
	jwt *jwt.TokenUser
	hash *utils.Hash
	cache bigcache.Cache
	snowflake *snowflake.Snowflake
	db *gorm.DB
}

func NewServerUser(server *gin.Engine,config *config.Config, jwt *jwt.TokenUser, hash *utils.Hash, cache bigcache.Cache, snowflake *snowflake.Snowflake,db *gorm.DB) *UserServer{
	return &UserServer{
		server:server,
		config:config,
		jwt:jwt,
		hash:hash,
		cache:cache,
		snowflake:snowflake,
		db:db,
	}
}

func (userServer *UserServer) Run(){
	repo :=repository.NewUserRepository(userServer.db)
	useCase:= usecase.NewUserCase(repo,userServer.hash,userServer.jwt,userServer.config)
	userHTTP := http.NewUserHTTP(useCase,userServer.snowflake)
	delivery.NewUserHTTP(userServer.server,userServer.cache,userHTTP)
}