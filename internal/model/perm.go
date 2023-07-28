package model

import gormadapter "github.com/casbin/gorm-adapter/v3"

type PermRole struct {
	IIDModel
	Name        string `gorm:"type:varchar(255);not null;default:''"`
	Description string `gorm:"type:varchar(512);not null;default:''"`
}

type PermResource struct {
	IIDModel
	Name        string `gorm:"type:varchar(255);not null;default:''"`
	Description string `gorm:"type:varchar(512);not null;default:''"`
}

type PermApi struct {
	IIDModel
	Path        string `gorm:"uniqueIndex"`
	ResourceID  uint64
	Description string `gorm:"type:varchar(512);not null;default:''"`
	Action      string `gorm:"type:varchar(255);not null;default:''"`
	Resource    *PermResource
}

type PermAccountRole struct {
	AccountID string `gorm:"type:varchar(255);not null;default:'';primarykey"`
	RoleID    uint64 `gorm:"type:varchar(255);not null;default:'';primarykey"`
	Time
}

// p : role_id, resource_id, action

type CasbinRule struct {
	gormadapter.CasbinRule
}

type PermRoleReference struct {
	Casbin  int64 `json:"casbin"`
	Account int64 `json:"account"`
}

type PermResourceReference struct {
	Casbin int64 `json:"casbin"`
	API    int64 `json:"api"`
}
