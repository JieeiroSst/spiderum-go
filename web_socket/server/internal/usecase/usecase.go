package usecase

import (
	"github.com/gorilla/websocket"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/model"
	"gitlab.com/Spide_IT/spide_it/web_socket/pkg/snowflake"
	"net/http"
)

type WebUsecase struct {
	repo internal.WebSocketRepository
	client *model.Client
	pool *model.Pool
	snowflake snowflake.Snowflake
	upgrade websocket.Upgrader
}

func NewWebUsecase(repo internal.WebSocketRepository,pool *model.Pool, client *model.Client, snowflake snowflake.Snowflake, upgrade websocket.Upgrader) *WebUsecase{
	return &WebUsecase{
		repo:repo,
		client:client,
		snowflake:snowflake,
		upgrade:upgrade,
		pool:pool,
	}
}

func (w *WebUsecase) Read(){
	defer func() {
		w.client.Pool.Unregister <- w.client
		w.client.Conn.Close()
	}()
	for {
		messageType, body ,err:=w.client.Conn.ReadMessage()
		if err!=nil{
			return
		}
		userId:=w.repo.GetByIdUser(messageType)
		message:=model.Message{
			Type:   messageType,
			Body:   string(body),
			UserId: userId,
		}
		w.client.Pool.Broadcast <- message
	}
}

func (w *WebUsecase) Start(name string){
	for {
		select {
			case client:=<-w.pool.Register:
				w.pool.Clients[client] = true
			for client, _ := range w.pool.Clients {
				id:=w.snowflake.GearedID()
				var user = model.User{
					Id: w.snowflake.GearedID(),
					Name: name,
				}
				err:=w.repo.CreateUser(user)
				if err!=nil{
					break
				}
				message:=model.Message{Type:id,Body:"New User Joined...",UserId:id}
				_ =w.repo.CreateMessage(message)
				client.Conn.WriteJSON(message)
			}
			break
			case client:=<-w.pool.Unregister:
				delete(w.pool.Clients, client)
				for client, _ := range w.pool.Clients {
					id:=w.snowflake.GearedID()
					UserId:=w.repo.GetByNameUser(name)
					message:=model.Message{Type: id, Body: "User Disconnected...",UserId:UserId}
					_ =w.repo.CreateMessage(message)
					client.Conn.WriteJSON(message)
				}
			break
			case message:=<-w.pool.Broadcast:
				for client,_:=range w.pool.Clients {
					err:=client.Conn.WriteJSON(message)
					if err!=nil{
						return
					}
					UserId:=w.repo.GetByNameUser(name)
					message.UserId=UserId
					message.Type=w.snowflake.GearedID()
					err = w.repo.CreateMessage(message)
					if err!=nil{
						return
					}
				}
		}
	}
}

func (w *WebUsecase) Upgrade(response http.ResponseWriter, request *http.Request) (*websocket.Conn, error){
	conn,err :=w.upgrade.Upgrade(response,request,nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}