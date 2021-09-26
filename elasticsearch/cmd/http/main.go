package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/config"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/elasticsearch/server"
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
		// Handle error
		panic(err)
	}
	log.Printf("Elasticsearch version %s %s\n", esversion,">>>>")

	exists, err := elasticsearchConn.IndexExists("posts").Do(context.Background())
	if err != nil {
		// Handle error
	}
	log.Println(exists,">....................")
	if !exists {
		// Index does not exist yet.
	}

	httpServer:=server.NewElasticsearchServer(router,elasticsearchConn)
	httpServer.Run()
	router.Run(config.Server.PortServerHttp)
}