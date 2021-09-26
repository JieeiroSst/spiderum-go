package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	deliveryhttp "gitlab.com/Spide_IT/spide_it/elasticsearch/internal/elasticsearch/delivery/http"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/elasticsearch/http"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/elasticsearch/repository"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/elasticsearch/usecase"
)

type ElasticsearchServer struct {
	server *gin.Engine
	elastic *elastic.Client
}

func NewElasticsearchServer(server *gin.Engine,elastic *elastic.Client)*ElasticsearchServer{
	return &ElasticsearchServer{
		server:server,
		elastic:elastic,
	}
}

func (server *ElasticsearchServer) Run(){
	repo:=repository.NewElasticsearchInterface(server.elastic)
	usecase:=usecase.NewElasticsearchUsecase(repo)
	http:=http.NewHttp(usecase)
	deliveryhttp.NewHttp(server.server,http)
}