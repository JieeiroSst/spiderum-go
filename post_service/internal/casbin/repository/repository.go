package repository

import (
	"fmt"
	"gitlab.com/Spide_IT/spide_it/internal/casbin/model"
	"gorm.io/gorm"
)

type CasbinRuleRepository struct {
	db *gorm.DB
}

func NewCasbinRuleRepository(db *gorm.DB) *CasbinRuleRepository{
	return &CasbinRuleRepository{
		db:db,
	}
}

func (repo *CasbinRuleRepository) CasbinRuleAll() ([]model.CasbinRule,error){
	var casbinRules []model.CasbinRule
	repo.db.Table("casbin_rule").Select("*").Scan(&casbinRules)
	return casbinRules, nil
}

func (repo *CasbinRuleRepository) CasbinRuleById(id int) (model.CasbinRule,error){
	var casbinRule model.CasbinRule
	repo.db.Table("casbin_rule").Select("*").Where("id = ?",id).Scan(&casbinRule)
	return casbinRule, nil
}

func (repo *CasbinRuleRepository)  UpdateCasbinRulePtype(id int,ptype string) error{
	err:=repo.db.Exec("UPDATE `casbin_rule` SET ptype = ?  WHERE id = ?", ptype, id)
	return err.Error
}

func (repo *CasbinRuleRepository)  UpdateCasbinRuleName(id int,name string) error{
	err:=repo.db.Exec("UPDATE `casbin_rule` SET v0 = ?  WHERE id = ?", name, id)
	return err.Error
}
func (repo *CasbinRuleRepository)  UpdateCasbinRuleEndpoint(id int,endpoint string) error{
	err:=repo.db.Exec("UPDATE `casbin_rule` SET v1 = ? WHERE id = ? ", endpoint, id)
	return err.Error
}
func (repo *CasbinRuleRepository)  UpdateCasbinMethod(id int,method string) error{
	err:=repo.db.Exec("UPDATE `casbin_rule` SET v2 = ?  WHERE id = ?", method, id)
	return err.Error
}

func (repo *CasbinRuleRepository) CreateCasbinRule(casbin model.CasbinRule) error{
	stmtString := fmt.Sprintf("INSERT INTO `casbin_rule` (ptype,v0,v1,v2) VALUES ('%s','%s','%s','%s');", casbin.Ptype,casbin.V0,casbin.V1,casbin.V2)
	err:=repo.db.Exec(stmtString)
	return err.Error
}

func (repo *CasbinRuleRepository) DeleteCasbinRule(id int) error{
	err:=repo.db.Exec("DELETE FROM `casbin_rule` where id = ? ", id)
	return err.Error
}