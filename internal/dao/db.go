package dao

import (
	"github.com/cro4k/authorize/config"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	var dialer gorm.Dialector
	switch config.C().DB.Driver {
	case "", config.MYSQL:
		dialer = mysql.New(mysql.Config{DSN: config.C().DB.DSN})
	case config.SQLITE:
		dialer = sqlite.Open(config.C().DB.DSN)
	}

	var err error
	db, err = gorm.Open(
		dialer,
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}
}
