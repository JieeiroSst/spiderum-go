package ip

import (
	"gitlab.com/Spide_IT/spide_it/internal/ip/model"
)

type IpUsecase interface {
	FindIpAll() ([]model.Ip,error)
}