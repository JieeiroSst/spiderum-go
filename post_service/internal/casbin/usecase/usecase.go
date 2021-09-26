package usecase

import (
	"gitlab.com/Spide_IT/spide_it/internal/casbin"
	"gitlab.com/Spide_IT/spide_it/internal/casbin/model"
)

type CasbinRuleUseCase struct {
	repo casbin.CasbinRuleRepository
}

func NewCasbinRuleUseCase(repo casbin.CasbinRuleRepository) *CasbinRuleUseCase {
	return &CasbinRuleUseCase{
		repo:repo,
	}
}


func(repo *CasbinRuleUseCase) CasbinRuleAll() ([]model.CasbinRule,error){
	casbin,err:=repo.repo.CasbinRuleAll()
	return casbin, err
}

func(repo *CasbinRuleUseCase) CasbinRuleById(id int) (model.CasbinRule,error){
	casbin,err:=repo.repo.CasbinRuleById(id)
	return casbin,err
}

func(repo *CasbinRuleUseCase) CreateCasbinRule(casbin model.CasbinRule) error{
	return repo.repo.CreateCasbinRule(casbin)
}

func(repo *CasbinRuleUseCase) DeleteCasbinRule(id int) error{
	return repo.repo.DeleteCasbinRule(id)
}

func(repo *CasbinRuleUseCase) UpdateCasbinRulePtype(id int,ptype string) error{
	return repo.repo.UpdateCasbinRulePtype(id,ptype)
}

func(repo *CasbinRuleUseCase) UpdateCasbinRuleName(id int,name string) error{
	return repo.repo.UpdateCasbinRuleName(id,name)
}

func(repo *CasbinRuleUseCase) UpdateCasbinRuleEndpoint(id int,endpoint string) error{
	return repo.repo.UpdateCasbinRuleEndpoint(id,endpoint)
}

func(repo *CasbinRuleUseCase) UpdateCasbinMethod(id int,method string) error{
	return repo.repo.UpdateCasbinMethod(id,method)
}