package repository

import (
	"gitlab.com/Spide_IT/spide_it/email/internal/model"
	"gitlab.com/Spide_IT/spide_it/email/pkg/snowflake"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type EmailRepository struct {
	snowflake snowflake.Snowflake
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB,snowflake snowflake.Snowflake) *EmailRepository {
	_ = db.AutoMigrate(&model.Email{})
	return &EmailRepository{db:db,snowflake:snowflake}
}

func (e *EmailRepository) CreateEmail(el model.Email) (create model.Email,err error) {
	create = model.Email{
		Id:     strconv.Itoa(e.snowflake.GearedID()),
		Email:       el.Email,
		To:          el.To,
		From:        el.From,
		Body:        el.Body,
		Subject:     el.Subject,
		ContentType: el.ContentType,
		CreatedAt:   time.Now(),
	}
	err = e.db.Create(&create).Error
	return
}

