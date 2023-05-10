package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Open(c Config) (*gorm.DB, error) {
	return gorm.Open(c.dialector(), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
}

func MustOpen(c Config) *gorm.DB {
	if db, err := Open(c); err != nil {
		panic(err)
	} else {
		return db
	}
}

var db *gorm.DB

func Setup(c Config) {
	db = MustOpen(c)
}

func DB() *gorm.DB {
	return db
}

func TX(f func(tx *gorm.DB) error) error {
	return db.Transaction(f)
}
