package repository

import (
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/user/model"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	if err:=db.AutoMigrate(&model.Users{});err!=nil {
		log.Println(err)
	}
	return &UserRepository{db:db}
}

func (d *UserRepository) CheckAccount(user model.Users) (int,string,bool) {
	var result model.Users
	r := d.db.Where("username = ? ", user.Username).Limit(1).Find(&result)

	if r.Error != nil{
		return -1,"",false
	}

	if result.Id == 0 {
		return -1,"",false
	}

	return result.Id,result.Password,true
}

func (d *UserRepository) CheckAccountExists(user model.Users) bool {
	var result model.Users
	r := d.db.Where("username = ? ", user.Username).Limit(1).Find(&result)
	if r.Error != nil{
		return false
	}

	if result.Id !=0 {
		return false
	}

	return true
}

func (d *UserRepository) CreateAccount(user model.Users) error {
	return d.db.Create(&user).Error
}

func (d *UserRepository) RequestIpComputer(ip ip.Ip) error{
	return d.db.Create(&ip).Error
}