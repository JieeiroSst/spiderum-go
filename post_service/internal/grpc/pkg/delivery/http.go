package delivery

import (
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg"
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/api"
)

type Grpc struct {
	usecase pkg.GrpcUsecase
}

func NewGrpc(usecase pkg.GrpcUsecase) *Grpc{
	return &Grpc{usecase:usecase}
}

func (grpc *Grpc) GetData() ([]*api.Post,error){
	posts,err:=grpc.usecase.GetData()
	if err!=nil{
		return nil, err
	}
	return posts,nil
}