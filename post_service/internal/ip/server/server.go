package server

import (
	"github.com/gin-gonic/gin"
	delivery "gitlab.com/Spide_IT/spide_it/internal/ip/delivery/http"
	"gitlab.com/Spide_IT/spide_it/internal/ip/http"
	"gitlab.com/Spide_IT/spide_it/internal/ip/repository"
	"gitlab.com/Spide_IT/spide_it/internal/ip/usecase"
	auth "gitlab.com/Spide_IT/spide_it/middleware"
	"gitlab.com/Spide_IT/spide_it/pkg/bigcache"
	"gitlab.com/Spide_IT/spide_it/pkg/jwt"
	"gitlab.com/Spide_IT/spide_it/pkg/snowflake"
	"gorm.io/gorm"
)

type IpServer struct {
	server *gin.Engine
	db *gorm.DB
	snowflake *snowflake.Snowflake
	cache bigcache.Cache
	jwt *jwt.TokenUser
}

func NewIpServer(server *gin.Engine, db *gorm.DB, snowflake *snowflake.Snowflake,cache bigcache.Cache,jwt *jwt.TokenUser) *IpServer{
	return &IpServer{
		server:server,
		db:db,
		snowflake:snowflake,
		cache:cache,
		jwt:jwt,
	}
}

func (ipserver *IpServer) Run(){
	repo:=repository.NewIpRepository(ipserver.db)
	usecase:=usecase.NewIpUsecase(repo)
	http:=http.Newhttp(usecase,*ipserver.snowflake)
	auth:=auth.NewAuthorization(ipserver.cache,ipserver.jwt)
	delivery.NewIpHttp(ipserver.server,*http,*auth,ipserver.db)
}