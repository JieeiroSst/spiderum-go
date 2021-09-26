package usecase

import (
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg"
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/api"
)

type GrpcUsecase struct {
	repo pkg.GrpcRepository
}

func NewGrpcUsecase(repo pkg.GrpcRepository) *GrpcUsecase{
	return &GrpcUsecase{repo:repo}
}

func (grpc *GrpcUsecase) GetData() ([]*api.Post,error){
	posts,err:=grpc.repo.GetData()
	if err!=nil{
		return nil, err
	}
	return posts,nil
}