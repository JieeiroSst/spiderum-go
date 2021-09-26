package usecase

import (
	"gitlab.com/Spide_IT/spide_it/config"
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/user"
	"gitlab.com/Spide_IT/spide_it/internal/user/model"
	"gitlab.com/Spide_IT/spide_it/pkg/jwt"
	"gitlab.com/Spide_IT/spide_it/utils"
	"log"
)

type UserUseCase struct {
	userRepo user.UserRepository
	hash *utils.Hash
	jwt *jwt.TokenUser
	conf *config.Config
}

func NewUserCase(userRepo user.UserRepository, hash *utils.Hash, jwt *jwt.TokenUser, conf *config.Config) *UserUseCase {
	return &UserUseCase{
		userRepo:userRepo,
		hash:hash,
		jwt:jwt,
		conf:conf,
	}
}

func (u *UserUseCase) Login(user model.Users) (string,int){
	id,hashPassword,check  := u.userRepo.CheckAccount(user)
	if check == false {
		return "User does not exist",id
	}
	if checkPass := u.hash.CheckPassowrd(user.Password, hashPassword); checkPass != nil {
		return "password entered incorrectly",-1
	}
	token, _ := u.jwt.GenerateToken(user.Username)
	return token,id
}

func (u *UserUseCase) SignUp(user model.Users) string{
	check := u.userRepo.CheckAccountExists(user)
	if check == false {
		return "user already exists"
	}
	hashPassword, err := u.hash.HashPassword(user.Password)
	if err != nil {
		log.Println("error server", err)
	}
	account:= model.Users{
		Id:           user.Id,
		Username:     user.Username,
		Password:     hashPassword,
	}
	err = u.userRepo.CreateAccount(account)
	if err!=nil{
		return "Create failed"
	}
	return user.Username+":"+"Create success"
}

func (u *UserUseCase) RequestIpComputer(ip ip.Ip) error{
	return u.userRepo.RequestIpComputer(ip)
}