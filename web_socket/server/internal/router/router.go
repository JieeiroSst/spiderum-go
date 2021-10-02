package router

import (
	"gitlab.com/Spide_IT/spide_it/web_socket/internal"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/model"
	"math/rand"
	"net/http"
)

type WebsocketRouter struct {
	usecase internal.WebUsecase
	server  internal.WebsocketServer
	pool   model.Pool
}

func NewWebsocketRouter(usecase internal.WebUsecase, server  internal.WebsocketServer, pool   model.Pool)*WebsocketRouter{
	return &WebsocketRouter{
		usecase:usecase,
		server:server,
		pool:pool,
	}
}

func (w *WebsocketRouter) SetupRouter(){
	bytes := make([]byte, 10)
	for i := 0; i < 10; i++ {
		bytes[i] = byte(rand.Intn(1-100))
	}
	go w.usecase.Start(string(bytes))

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		w.server.ServeWebsocket(&w.pool,writer,request)
	})
}

