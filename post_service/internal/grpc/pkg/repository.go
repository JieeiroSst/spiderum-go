package pkg

import (
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/api"
)

type GrpcRepository interface {
	GetData() ([]*api.Post,error)
}