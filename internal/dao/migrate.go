package dao

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/cro4k/authorize/internal/model"
)

var migrations = []*gormigrate.Migration{
	//{
	//	ID: "fix",
	//	Migrate: func(db *gorm.DB) error {
	//		return migration.With(db).
	//			CreateTable("example_table_1").
	//			AddColumn("example_table_2", "example_column").
	//			Error()
	//	},
	//	Rollback: func(db *gorm.DB) error {
	//		return migration.With(db).
	//			DropTable("example_table_1").
	//			DropColumn("example_table_2", "example_column").
	//			Error()
	//	},
	//},
}

func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	m.InitSchema(func(db *gorm.DB) error {
		return db.AutoMigrate(
			&model.Account{},
			&model.AccountProfile{},
			&model.AccountLog{},
			&model.Application{},
		)
	})
	if err := m.Migrate(); err != nil {
		logrus.Error(err)
	}
}
