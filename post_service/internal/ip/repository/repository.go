package repository

import (
	"gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gorm.io/gorm"
	"log"
)

type IpRepository struct {
	db *gorm.DB
}

func NewIpRepository(db *gorm.DB) *IpRepository{
	if err:=db.AutoMigrate(&model.Ip{});err!=nil {
		log.Println(err)
	}
	return &IpRepository{db:db}
}

func (i *IpRepository) FindIpAll() ([]model.Ip,error){
	var ips []model.Ip
	i.db.Find(&ips)
	return ips,nil
}