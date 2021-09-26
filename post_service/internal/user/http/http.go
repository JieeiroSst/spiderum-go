package http

import (
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/user"
	"gitlab.com/Spide_IT/spide_it/internal/user/model"
	"gitlab.com/Spide_IT/spide_it/pkg/snowflake"
	"time"
)

type UserHTTP struct {
	userCase user.UserUserCase
	snowflake *snowflake.Snowflake
}

func NewUserHTTP(userCase user.UserUserCase,snowflake *snowflake.Snowflake) *UserHTTP {
	return &UserHTTP{
		userCase:userCase,
		snowflake:snowflake,
	}
}

func (u *UserHTTP) Login(username,password string) (string,int){
	users := model.Users{
		Username: username,
		Password: password,
	}

	token,id:=u.userCase.Login(users)

	return token,id

}

func (u *UserHTTP) Signup(username,password string)string {
	users := model.Users{
		Id:       u.snowflake.GearedID(),
		Username: username,
		Password: password,
	}
	return u.userCase.SignUp(users)
}

func (u *UserHTTP) RequestIpComputer(ipAddress string,method string) error{
	ipMacAddress:=ip.Ip{
		Id:     u.snowflake.GearedID(),
		Ip:     ipAddress,
		Method: method,
		RequestAt:time.Now(),
	}
	return u.userCase.RequestIpComputer(ipMacAddress)
}