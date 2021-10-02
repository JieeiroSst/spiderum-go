package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocket struct {
	socket websocket.Upgrader
}

func NewWebSocket(socket websocket.Upgrader) *WebSocket{
	socket = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	return &WebSocket{socket:socket}
}
