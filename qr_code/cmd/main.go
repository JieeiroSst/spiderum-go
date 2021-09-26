package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/qr_code/config"
	qrcodeserver "gitlab.com/Spide_IT/spide_it/qr_code/internal/server"
	"log"
	"time"
)

func main(){
	log.Println("Starting server")
	router := gin.Default()

	routerConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(routerConfig))

	config, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Fatal(err)
	}

	qrcodeServer :=qrcodeserver.NewQrcodeServer(router)
	qrcodeServer.Run()

	if err := router.Run(config.Server.PortServer);err!=nil{
		log.Println(err)
	}
}