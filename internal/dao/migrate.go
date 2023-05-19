package dao

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/cro4k/authorize/internal/dao/migration"
	"github.com/cro4k/authorize/internal/model"
)

var migrations = []*gormigrate.Migration{
	{
		ID: "authorize_init_schema",
		Migrate: func(db *gorm.DB) error {
			return migration.With(db).CreateTable(
				&model.Account{},
				&model.AccountProfile{},
				&model.AccountLog{},
				&model.Application{},
				&model.PermRole{},
				&model.PermResource{},
				&model.PermApi{},
				&model.PermAccountRole{},
			).Error()
		},
		Rollback: func(db *gorm.DB) error {
			return migration.With(db).DropTable(
				&model.Account{},
				&model.AccountProfile{},
				&model.AccountLog{},
				&model.Application{},
				&model.PermRole{},
				&model.PermResource{},
				&model.PermApi{},
				&model.PermAccountRole{},
			).Error()
		},
	},
}

func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		logrus.Error(err)
	}
}
