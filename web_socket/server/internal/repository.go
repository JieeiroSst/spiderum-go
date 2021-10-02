package internal

import "gitlab.com/Spide_IT/spide_it/web_socket/internal/model"


type WebSocketRepository interface {
	CreateUser(user model.User) error
	CreateMessage(message model.Message) error
	GetByIdUser(idMessage int) int
	CheckExistsUser(name string) bool
	GetByNameUser(name string) int
}