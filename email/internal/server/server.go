package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"gitlab.com/Spide_IT/spide_it/email/config"
	"gitlab.com/Spide_IT/spide_it/email/internal/delivery/http"
	"gitlab.com/Spide_IT/spide_it/email/internal/delivery/rabbitmq"
	"gitlab.com/Spide_IT/spide_it/email/internal/repository"
	"gitlab.com/Spide_IT/spide_it/email/internal/usecase"
	EmailPkg "gitlab.com/Spide_IT/spide_it/email/pkg/email"
	"gitlab.com/Spide_IT/spide_it/email/pkg/snowflake"
	"gitlab.com/Spide_IT/spide_it/email/utils"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type EmailServer struct {
	server 	   *gin.Engine
	db         *gorm.DB
	amqpConn   *amqp.Connection
	cfg        *config.Config
	snowflake  *snowflake.Snowflake
	utils      *utils.EmailUtils
	emailPkg   *EmailPkg.ServicePkg
}


func NewEmailsServer(server *gin.Engine,amqpConn *amqp.Connection, cfg *config.Config, db *gorm.DB,snowflake *snowflake.Snowflake,utils *utils.EmailUtils,EmailPkg *EmailPkg.ServicePkg) *EmailServer {
	return &EmailServer{
		server:server,
		amqpConn: amqpConn,
		cfg: cfg,
		db: db,
		utils:utils,
		emailPkg:EmailPkg,
		snowflake:snowflake,
	}
}

func (e *EmailServer) Run() {
	emailsPublisher, err := rabbitmq.NewEmailsPublisher(e.cfg)
	if err != nil {
		log.Println("Emails Publisher can't initialized")
	}
	defer emailsPublisher.CloseChan()

	err = emailsPublisher.SetupExchangeAndQueue(
		e.cfg.RabbitMQ.Exchange,
		e.cfg.RabbitMQ.Queue,
		e.cfg.RabbitMQ.RoutingKey,
		e.cfg.RabbitMQ.ConsumerTag, )
	if err!=nil{
		log.Println(err)
	}

	repo:=repository.NewEmailRepository(e.db,*e.snowflake)
	usecase:=usecase.NewEmailUseCase(repo,emailsPublisher,*e.utils,*e.emailPkg)
	http.NewHttp(e.server,usecase)
	emailsAmqpConsumer := rabbitmq.NewImagesConsumer(e.amqpConn, usecase,e.cfg)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := emailsAmqpConsumer.StartConsumer(
			e.cfg.RabbitMQ.WorkerPoolSize,
			e.cfg.RabbitMQ.Exchange,
			e.cfg.RabbitMQ.Queue,
			e.cfg.RabbitMQ.RoutingKey,
			e.cfg.RabbitMQ.ConsumerTag,
		)
		if err != nil {
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		log.Printf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Printf("ctx.Done: %v", done)
	}
}