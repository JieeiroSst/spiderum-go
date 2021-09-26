package casbin

import "gitlab.com/Spide_IT/spide_it/internal/casbin/model"

type CasbinRuleUsecase interface {
	CasbinRuleAll() ([]model.CasbinRule,error)
	CasbinRuleById(id int) (model.CasbinRule,error)
	CreateCasbinRule(casbin model.CasbinRule) error
	DeleteCasbinRule(id int) error
	UpdateCasbinRulePtype(id int,ptype string) error
	UpdateCasbinRuleName(id int,name string) error
	UpdateCasbinRuleEndpoint(id int,endpoint string) error
	UpdateCasbinMethod(id int,method string) error
}