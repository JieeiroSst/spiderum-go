package http

import(
	"gitlab.com/Spide_IT/spide_it/internal/casbin"
	"gitlab.com/Spide_IT/spide_it/internal/casbin/model"
)

type Http struct {
	usecase casbin.CasbinRuleUsecase
}

func NewHttp(usecase casbin.CasbinRuleUsecase) *Http{
	return &Http{
		usecase:usecase,
	}
}

func(http *Http) CasbinRuleAll() ([]model.CasbinRule,error){
	casbin,err:=http.usecase.CasbinRuleAll()
	return casbin, err
}

func(http *Http) CasbinRuleById(id int) (model.CasbinRule,error){
	casbin,err:=http.usecase.CasbinRuleById(id)
	return casbin,err
}

func(http *Http) CreateCasbinRule(casbin model.CasbinRule) error{
	return http.usecase.CreateCasbinRule(casbin)
}

func(http *Http) DeleteCasbinRule(id int) error{
	return http.usecase.DeleteCasbinRule(id)
}

func(http *Http) UpdateCasbinRulePtype(id int,ptype string) error{
	return http.usecase.UpdateCasbinRulePtype(id,ptype)
}

func(http *Http) UpdateCasbinRuleName(id int,name string) error{
	return http.usecase.UpdateCasbinRuleName(id,name)
}

func(http *Http) UpdateCasbinRuleEndpoint(id int,endpoint string) error{
	return http.usecase.UpdateCasbinRuleEndpoint(id,endpoint)
}

func(http *Http) UpdateCasbinMethod(id int,method string) error{
	return http.usecase.UpdateCasbinMethod(id,method)
}