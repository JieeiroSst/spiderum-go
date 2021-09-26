package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/upload/config"
	"gitlab.com/Spide_IT/spide_it/upload/internal/s3/server"
	"gitlab.com/Spide_IT/spide_it/upload/pkg/s3"
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

	s3:=s3.NewS3(config.S3.S3Region)

	server:=server.NewUploadServer(router,s3)
	server.Run()

	router.Run(config.Server.PortServer)
}