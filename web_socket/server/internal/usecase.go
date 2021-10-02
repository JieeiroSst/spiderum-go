package internal

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WebUsecase interface {
	Read()
	Start(name string)
	Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
}