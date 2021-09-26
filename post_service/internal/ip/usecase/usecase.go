package usecase

import (
	"gitlab.com/Spide_IT/spide_it/internal/ip"
	"gitlab.com/Spide_IT/spide_it/internal/ip/model"
)

type IpUsecase struct {
	repo ip.IpRepository
}

func NewIpUsecase(repo ip.IpRepository) *IpUsecase{
	return &IpUsecase{repo:repo}
}


func (i *IpUsecase) FindIpAll() ([]model.Ip,error){
	return i.repo.FindIpAll()
}