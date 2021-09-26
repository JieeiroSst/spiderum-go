package server

import (
	"github.com/gin-gonic/gin"
	delivery "gitlab.com/Spide_IT/spide_it/internal/post/delivery/http"
	"gitlab.com/Spide_IT/spide_it/internal/post/http"
	"gitlab.com/Spide_IT/spide_it/internal/post/repository"
	"gitlab.com/Spide_IT/spide_it/internal/post/usecase"
	auth "gitlab.com/Spide_IT/spide_it/middleware"
	"gitlab.com/Spide_IT/spide_it/pkg/bigcache"
	"gitlab.com/Spide_IT/spide_it/pkg/jwt"
	"gitlab.com/Spide_IT/spide_it/pkg/snowflake"
	"gorm.io/gorm"
)

type PostServer struct {
	server *gin.Engine
	db *gorm.DB
	snowflake *snowflake.Snowflake
	jwt *jwt.TokenUser
	cache bigcache.Cache
}

func NewPostServer(server *gin.Engine,db *gorm.DB,snowflake *snowflake.Snowflake,jwt *jwt.TokenUser,cache bigcache.Cache) *PostServer {
	return &PostServer{
		server:server,
		db:db,
		snowflake:snowflake,
		jwt:jwt,
		cache:cache,
	}
}

func (postServer *PostServer) Run(){
	repo:=repository.NewPostsRepository(postServer.db)
	usecase:=usecase.NewPostsUsecase(repo)
	http:=http.NewPostHttp(usecase,*postServer.snowflake)
	auth:=auth.NewAuthorization(postServer.cache,postServer.jwt)
	delivery.NewPostHttp(postServer.server,*http,*auth,postServer.db)
}