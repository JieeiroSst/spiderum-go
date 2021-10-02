package repository

import (
	"errors"
	"gitlab.com/Spide_IT/spide_it/web_socket/internal/model"
	"gorm.io/gorm"
)

type WebSocketRepository struct {
	db *gorm.DB
}

func NewWebSocketRepository(db *gorm.DB) *WebSocketRepository{
	return &WebSocketRepository{db:db}
}

func (w *WebSocketRepository) CheckExistsUser(name string) bool{
	var check int
	w.db.Raw("SELECT EXISTS(SELECT name FROM  users)").Scan(&check)
	if check == 0 {
		return false
	}
	return true
}

func (w *WebSocketRepository) CreateUser(user model.User) error {
	var check int
	w.db.Raw("SELECT EXISTS(SELECT name FROM  users)").Scan(&check)
	if check == 0 {
		errors.New("name exists table users")
	}
	err :=w.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (w *WebSocketRepository) CreateMessage(message model.Message) error{
	err:=w.db.Create(&message).Error
	if err != nil {
		return err
	}
	return nil
}

func (w *WebSocketRepository) GetByIdUser(idMessage int) int{
	var result struct {
		UserId int  `json:"user_id"`
	}
	w.db.Model(&model.Message{}).Select("user_id").Where("id = ?",idMessage).Find(&result)
	return result.UserId
}

func (w *WebSocketRepository)GetByNameUser(name string) int {
	var result struct {
		id int  `json:"id"`
	}
	w.db.Model(&model.Message{}).Select("id").Where("name = ?",name).Find(&result)
	return result.id
}