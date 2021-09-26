package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gitlab.com/Spide_IT/spide_it/config"
	casbinServer "gitlab.com/Spide_IT/spide_it/internal/casbin/server"
	postServer "gitlab.com/Spide_IT/spide_it/internal/post/server"
	userServer "gitlab.com/Spide_IT/spide_it/internal/user/server"
	ipServer "gitlab.com/Spide_IT/spide_it/internal/ip/server"
	"gitlab.com/Spide_IT/spide_it/pkg/bigcache"
	"gitlab.com/Spide_IT/spide_it/pkg/jwt"
	"gitlab.com/Spide_IT/spide_it/pkg/mysql"
	"gitlab.com/Spide_IT/spide_it/pkg/snowflake"
	"gitlab.com/Spide_IT/spide_it/utils"
	"log"
	"time"

	_ "gitlab.com/Spide_IT/spide_it/docs"


)

func main(){
	config, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Fatal(err)
	}

	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Mysql.MysqlUser,
		config.Mysql.MysqlPassword,
		config.Mysql.MysqlHost,
		config.Mysql.MysqlPort,
		config.Mysql.MysqlDbname,
	)

	mysqlOrm:= mysql.NewMysqlConn(dns)

	log.Println("Starting server")

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routerConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(routerConfig))

	var hash = utils.NewHash()
	var snowflake = snowflake.NewSnowflake()
	var cache = *bigcache.NewBigCache()


	tokenUser :=jwt.NewTokenUser(config)

	serverUser :=userServer.NewServerUser(router,config, tokenUser,hash,cache, snowflake,mysqlOrm)
	newCabinServer :=casbinServer.NewCasbinServer(router,mysqlOrm)
	postServer:=postServer.NewPostServer(router,mysqlOrm,snowflake, tokenUser,cache)
	ipServer:=ipServer.NewIpServer(router,mysqlOrm,snowflake,cache,tokenUser)

	serverUser.Run()
	newCabinServer.Run()
	postServer.Run()
	ipServer.Run()

	if err := router.Run(config.Server.PortServer);err!=nil{
		log.Println(err)
	}
}