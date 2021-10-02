package pool

import "gitlab.com/Spide_IT/spide_it/web_socket/internal/model"

func NewPool() *model.Pool {
	return &model.Pool{
		Register:   make(chan *model.Client),
		Unregister: make(chan *model.Client),
		Clients:    make(map[*model.Client]bool),
		Broadcast:  make(chan model.Message),
	}
}
