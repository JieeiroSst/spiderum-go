package internal

import (
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/model"
	"net/http"
)

type WebsocketServer interface {
	ServeWebsocket(pool *model.Pool, w http.ResponseWriter, r *http.Request)
}
