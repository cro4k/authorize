package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MySQL  = "mysql"
	SQLite = "sqlite"
)

type Config struct {
	Driver string `json:"driver" yaml:"driver"`
	DSN    string `json:"dsn"    yaml:"dsn"`
}

func (c Config) dialector() gorm.Dialector {
	switch c.Driver {
	case "mysql":
		return mysql.Open(c.DSN)
	case "sqlite":
		return sqlite.Open(c.DSN)
	default:
		panic("unsupported db driver: " + c.Driver)
	}
}
