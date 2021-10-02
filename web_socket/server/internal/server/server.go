package server

import (
	"gitlab.com/Spide_IT/spide_it/web_socket/internal"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/model"
	"net/http"
)

type WebsocketServer struct {
	usecase internal.WebUsecase
}

func NewWebsocketServer(usecase internal.WebUsecase)*WebsocketServer{
	return &WebsocketServer{usecase:usecase}
}

func (server *WebsocketServer) ServeWebsocket(pool *model.Pool, w http.ResponseWriter, r *http.Request) {
	conn,err:=server.usecase.Upgrade(w,r)
	if err!=nil{

	}
	client:=&model.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	server.usecase.Read()
}
