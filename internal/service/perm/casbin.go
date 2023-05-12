package perm

import (
	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

const (
	rbacModel = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act`
)

type CasbinService struct {
	e   *casbin.Enforcer
	adp *gormadapter.Adapter
}

func NewCasbinService(db *gorm.DB) (*CasbinService, error) {
	adp, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	m, err := model.NewModelFromString(rbacModel)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, adp)
	if err != nil {
		return nil, err
	}
	return &CasbinService{e: e, adp: adp}, err
}

func (s *CasbinService) Enforce(val ...interface{}) (bool, error) {
	return s.e.Enforce(val...)
}

func (s *CasbinService) AddPolicy(sec string, ptype string, rule []string) error {
	// note: sec is not used in gorm-adapter
	return s.adp.AddPolicy(sec, ptype, rule)
}

func (s *CasbinService) RemovePolicy(sec string, ptype string, rule []string) error {
	return s.adp.RemovePolicy(sec, ptype, rule)
}

func (s *CasbinService) AddPolicies(sec string, ptype string, rules [][]string) error {
	return s.adp.AddPolicies(sec, ptype, rules)
}

func (s *CasbinService) CasbinRules() ([]*gormadapter.CasbinRule, error) {
	var data []*gormadapter.CasbinRule
	err := s.adp.GetDb().Find(&data).Error
	return data, err
}
