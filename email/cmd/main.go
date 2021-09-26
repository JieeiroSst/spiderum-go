package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/email/config"
	emailServer "gitlab.com/Spide_IT/spide_it/email/internal/server"
	emailPkg "gitlab.com/Spide_IT/spide_it/email/pkg/email"
	"gitlab.com/Spide_IT/spide_it/email/pkg/postgres"
	"gitlab.com/Spide_IT/spide_it/email/pkg/rabbitmq"
	"gitlab.com/Spide_IT/spide_it/email/pkg/snowflake"
	"gitlab.com/Spide_IT/spide_it/email/utils"
	"log"
	"time"
)

func main(){
	router := gin.Default()

	routerConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(routerConfig))

	var snowflake = snowflake.NewSnowflake()
	var emailUtils = utils.NewEmailUtils()
	var emailPkg = emailPkg.NewServicePkg()


	config, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Fatal(err)
	}
	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Postgres.PostgresqlHost,config.Postgres.PostgresqlUser,config.Postgres.PostgresqlPassword,
		config.Postgres.PostgresqlDbname,config.Postgres.PostgresqlPort)
	postgresConn:= postgres.NewPostgresConn(dsn)

	amqpConn, err := rabbitmq.NewRabbitMQConn(config)
	if err != nil {
		log.Println(err)
	}
	defer amqpConn.Close()

	emailsServer := emailServer.NewEmailsServer(router,amqpConn,config,postgresConn,snowflake,emailUtils,emailPkg)


	emailsServer.Run()

	router.Run(config.Server.PortServer)
}