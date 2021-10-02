package main

import (
	"fmt"
	"gitlab.com/Spide_IT/spide_it/web_socket/config"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/model"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/repository"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/router"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/server"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/usecase"
	"gitlab.com/Spide_IT/spide_it/web_socket/pkg/mysql"
	"gitlab.com/Spide_IT/spide_it/web_socket/pkg/pool"
	"gitlab.com/Spide_IT/spide_it/web_socket/pkg/snowflake"
	"gitlab.com/Spide_IT/spide_it/web_socket/pkg/websocket"
	"log"
	"net/http"
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
	pool:=pool.NewPool()
	snowflake := snowflake.NewSnowflake()
	upgrade:=websocket.NewWebSocket()
	client:=model.Client{}

	log.Println("Starting server")

	repo:=repository.NewWebSocketRepository(mysqlOrm)
	usecase:=usecase.NewWebUsecase(repo,pool,&client,*snowflake,*upgrade)
	server:=server.NewWebsocketServer(usecase)
	router:=router.NewWebsocketRouter(usecase,server,*pool)

	router.SetupRouter()

	http.ListenAndServe(":8000", nil)
}
