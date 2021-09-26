package server

import (
	"context"
	"github.com/olivere/elastic/v7"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/config"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/http"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/proto"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/repository"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/usecase"
	"google.golang.org/grpc"
	"log"
)

type Timer struct {
	config *config.Config
	elastic *elastic.Client
}

func NewTimer(config *config.Config,elastic *elastic.Client)*Timer{
	return &Timer{
		config:config,
		elastic:elastic,
	}
}

func (w *Timer) Run(){
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
	req :=&proto.RequestPost{}
	data,err:=client.UpdatePost(context.Background(),req)
	if err!=nil{
		log.Println("get data failed")
	}
	for _,post:=range data.Posts {
		err := http.InsertPost(context.Background(),*post)
		if err!=nil {
			log.Println("insert failed")
		}
	}
}