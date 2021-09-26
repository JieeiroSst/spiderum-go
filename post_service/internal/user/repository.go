package user

import (
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/user/model"
)

type UserRepository interface {
	CheckAccount(user model.Users) (int,string,bool)
	CheckAccountExists(user model.Users) bool
	CreateAccount(user model.Users) error
	RequestIpComputer(ip ip.Ip) error
}