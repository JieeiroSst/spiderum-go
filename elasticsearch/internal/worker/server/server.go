package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/config"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/http"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/proto"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/repository"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/usecase"
	"google.golang.org/grpc"
	"log"
)

type WorkerServer struct {
	server *gin.Engine
	config *config.Config
	elastic *elastic.Client
}

func NewWorkerServer(server *gin.Engine,config *config.Config,elastic *elastic.Client)*WorkerServer{
	return &WorkerServer{
		server:server,
		config:config,
		elastic:elastic,
	}
}

func (w *WorkerServer) Run() {
	conn,err:=grpc.Dial("localhost"+w.config.Server.PprofPort,grpc.WithInsecure())
	if err != nil {
		log.Println("CLIENT IS NO DIAL",err)
		return
	}
	log.Printf("CLIENT IS DIAL AT %s...", w.config.Server.PprofPort)

	client:=proto.NewHandleServiceClient(conn)

	repo:=repository.NewElasticsearchInterface(w.elastic)
	usecase:=usecase.NewElasticsearchUsecase(repo)
	http:=http.NewHttp(usecase)

	group:= w.server.Group("/api")
	group.POST("/insert", func(context *gin.Context) {
		req :=&proto.RequestPost{}
		data,err:=client.UpdatePost(context,req)
		if err!=nil{
			context.JSON(500,"get data failed")
			return
		}
		for _,post:=range data.Posts {
			err := http.InsertPost(context,*post)
			if err!=nil {
				context.JSON(500,"insert failed")
				return
			}
		}
		context.JSON(200,"insert success")
	})
}