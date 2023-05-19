package perm

import (
	"net/http"

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

type Service struct {
	e   *casbin.Enforcer
	adp *gormadapter.Adapter
	res *resources
}

func NewService(db *gorm.DB) (*Service, error) {
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
	return &Service{e: e, adp: adp}, err
}

func (s *Service) Enforce(role string, path string, method string) (bool, error) {
	resource, err := s.res.Find(path)
	if err != nil {
		return false, err
	}
	return s.enforce(role, resource, method)
}

func (s *Service) enforce(val ...interface{}) (bool, error) {
	return s.e.Enforce(val...)
}

func (s *Service) AddPolicy(sec string, ptype string, rule []string) error {
	// note: sec is not used in gorm-adapter
	return s.adp.AddPolicy(sec, ptype, rule)
}

func (s *Service) RemovePolicy(sec string, ptype string, rule []string) error {
	return s.adp.RemovePolicy(sec, ptype, rule)
}

func (s *Service) AddPolicies(sec string, ptype string, rules [][]string) error {
	return s.adp.AddPolicies(sec, ptype, rules)
}

func (s *Service) CasbinRules() ([]*gormadapter.CasbinRule, error) {
	var data []*gormadapter.CasbinRule
	err := s.adp.GetDb().Find(&data).Error
	return data, err
}

func (s *Service) extract(r *http.Request) {

}
