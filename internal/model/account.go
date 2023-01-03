package model

import "github.com/cro4k/authorize/internal/model/resource"

type Account struct {
	UIDModel
	Username          string     `gorm:"type:varchar(32);not null;uniqueIndex"`
	Password          string     `gorm:"type:varchar(255);not null"`
	Secret            string     `gorm:"type:varchar(255);not null"`
	Nonce             string     `gorm:"type:varchar(255);not null"`
	Email             CipherText `gorm:"type:varchar(64);not null;default:''"`
	EmailHash         string     `gorm:"type:varchar(64);not null;uniqueIndex;default:''"`
	Cellphone         CipherText `gorm:"type:varchar(64);not null;default:''"`
	CellphoneHash     string     `gorm:"type:varchar(64);not null;uniqueIndex;default:''"`
	Status            int8       `gorm:"not null;default:0"`
	CertificateStatus int8       `gorm:"not null;default:0"`
	CertificateID     string     `gorm:"type:varchar(64);not null;default:''"`
}

type AccountLog struct {
	UIDModel
	AccountID string `gorm:"type:varchar(64);not null"`
	IP        string `gorm:"type:varchar(255);not null;default:''"`
	UA        string `gorm:"type:varchar(255);not null;default:''"`
	RequestID string `gorm:"type:varchar(128);not null;default:''"`
	ClientID  string `gorm:"type:varchar(128);not null;default:''"`
	LoginType int8   `gorm:"not null;default:0"`
}

type AccountProfile struct {
	UIDModel
	Nickname string          `gorm:"type:varchar(128);not null;default:''"`
	Avatar   resource.Avatar `gorm:"type:varchar(500);not null;default:''"`
	Gender   int8            `gorm:"not null;default:0"`
	Bio      string          `gorm:"type:varchar(500);not null;default:''"`
}
