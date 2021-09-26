package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/config"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/server"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/pkg/elasticsearch"
	"log"
	"time"
)

func main(){
	config, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	routerConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(routerConfig))

	elasticsearchConn :=elasticsearch.NewGetElasticsearchConn(config.Elasticsearch.Dns)

	esversion, err := elasticsearchConn.ElasticsearchVersion(config.Elasticsearch.Dns)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Elasticsearch version %s\n", esversion)


	workerServer:=server.NewWorkerServer(router,config,elasticsearchConn)
	workerServer.Run()

	router.Run(config.Server.PortServerWorker)
}