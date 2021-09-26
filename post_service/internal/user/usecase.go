package user

import (
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/user/model"
)

type UserUserCase interface {
	Login(user model.Users) (string,int)
	SignUp(user model.Users) string
	RequestIpComputer(ip ip.Ip) error
}
