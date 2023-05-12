package dao

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/cro4k/authorize/internal/model"
)

type PermAccess interface {
	CreateRole(role *model.PermRole) (uint64, error)
	DeleteRole(id uint64) error
	CheckRoleReference(id uint64) (*model.PermRoleReference, error)
	CreateResource(resource *model.PermResource) (uint64, error)
	DeleteResource(id uint64) error
	CheckResourceReference(id uint64) (*model.PermResourceReference, error)
	CreateAPI(api *model.PermApi) (uint64, error)
	DeleteAPI(id uint64) error
	AddRole(accountID string, roleID uint64) error
	RemoveRole(accountID string, roleID uint64) error
	RoleList() ([]*model.PermRole, error)
	ResourceList() ([]*model.PermResource, error)
	APIList() ([]*model.PermApi, error)
}

type perm struct {
	db *gorm.DB
}

func (p *perm) CreateRole(role *model.PermRole) (uint64, error) {
	var err error
	if role.ID == 0 {
		err = p.db.Create(role).Error
	} else {
		err = p.db.Save(role).Error
	}
	return role.ID, err
}

func (p *perm) DeleteRole(id uint64) error {
	reference, err := p.CheckRoleReference(id)
	if err != nil {
		return err
	}
	if reference.Casbin > 0 || reference.Account > 0 {
		return fmt.Errorf("role %d has been used in casbin_rule or perm_account_role", id)
	}

	return p.db.Delete(&model.PermRole{}, "id = ?", id).Error
}

func (p *perm) CheckRoleReference(id uint64) (*model.PermRoleReference, error) {
	var reference = new(model.PermRoleReference)
	str := strconv.Itoa(int(id))
	p.db.Table("casbin_rule").Where("v0 = ?", str).Count(&reference.Casbin)
	p.db.Model(&model.PermAccountRole{}).Where("role_id = ?", id).Count(&reference.Account)
	return reference, nil
}

func (p *perm) CreateResource(resource *model.PermResource) (uint64, error) {
	var err error
	if resource.ID == 0 {
		err = p.db.Create(resource).Error
	} else {
		err = p.db.Save(resource).Error
	}
	return resource.ID, err
}

func (p *perm) DeleteResource(id uint64) error {
	reference, err := p.CheckResourceReference(id)
	if err != nil {
		return err
	}
	if reference.Casbin > 0 || reference.API > 0 {
		return fmt.Errorf("resource %d has been used in casbin_rule or perm_api", id)
	}
	return p.db.Delete(&model.PermResource{}, "id = ?", id).Error
}

func (p *perm) CheckResourceReference(id uint64) (*model.PermResourceReference, error) {
	var reference = new(model.PermResourceReference)
	str := strconv.Itoa(int(id))
	p.db.Table("casbin_rule").Where("v1 = ?", str).Count(&reference.Casbin)
	p.db.Model(&model.PermApi{}).Where("resource_id = ?", id).Count(&reference.API)
	return reference, nil
}

func (p *perm) CreateAPI(api *model.PermApi) (uint64, error) {
	var err error
	if api.ID == 0 {
		err = p.db.Create(api).Error
	} else {
		err = p.db.Save(api).Error
	}
	return api.ID, err
}

func (p *perm) DeleteAPI(id uint64) error {
	return p.db.Delete(&model.PermApi{}, "id = ?", id).Error
}

func (p *perm) AddRole(accountID string, roleID uint64) error {
	return p.db.Create(&model.PermAccountRole{
		AccountID: accountID,
		RoleID:    roleID,
	}).Error
}

func (p *perm) RemoveRole(accountID string, roleID uint64) error {
	return p.db.Delete(&model.PermAccountRole{}, "account_id = ? AND role_id = ?", accountID, roleID).Error
}

func (p *perm) RoleList() ([]*model.PermRole, error) {
	var data []*model.PermRole
	err := p.db.Find(&data).Error
	return data, err
}
func (p *perm) ResourceList() ([]*model.PermResource, error) {
	var data []*model.PermResource
	err := p.db.Find(&data).Error
	return data, err
}
func (p *perm) APIList() ([]*model.PermApi, error) {
	var data []*model.PermApi
	err := p.db.Find(&data).Error
	return data, err
}

func Permission(db *gorm.DB) PermAccess {
	return &perm{db: db}
}
