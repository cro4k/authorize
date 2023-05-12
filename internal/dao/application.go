package dao

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	"gorm.io/gorm"

	"github.com/cro4k/authorize/internal/model"
)

type ApplicationAccess interface {
	oauth2.ClientStore
}

type application struct {
	db *gorm.DB
}

func (a *application) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	var app = new(model.Application)
	err := a.db.First(app, "id = ?", id).Error
	return app, err
}

func Application(db *gorm.DB) ApplicationAccess {
	return &application{db: db}
}
