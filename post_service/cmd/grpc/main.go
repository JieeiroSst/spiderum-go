package main

import (
	"fmt"
	"gitlab.com/Spide_IT/spide_it/config"
	"gitlab.com/Spide_IT/spide_it/pkg/mysql"
	"google.golang.org/grpc"
	"net"

	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/adder"
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/api"
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/delivery"
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/repository"
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/usecase"
	"log"
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
	repo:=repository.NewGrpcRepository(mysqlOrm)
	usecase:=usecase.NewGrpcUsecase(repo)
	http:=delivery.NewGrpc(usecase)

	s:=grpc.NewServer()
	srv:=&adder.GRPCServer{}
	srv.NewGRPCServer(http)
	api.RegisterHandleServiceServer(s,srv)
	l, err := net.Listen("tcp", config.Server.PprofPort)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}