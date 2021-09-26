package ip

import (
	"gitlab.com/Spide_IT/spide_it/internal/ip/model"
)


type IpRepository interface {
	FindIpAll() ([]model.Ip,error)
}