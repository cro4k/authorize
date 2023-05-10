package model

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/cro4k/common/crypto/aesutil"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UID struct {
	ID string `gorm:"type:varchar(128);PRIMARY_KEY;NOT NULL"`
}

func (u *UID) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

type IID struct {
	ID uint64 `gorm:"PRIMARY_KEY;NOT NULL;AUTO INCREMENT"`
}

type Time struct {
	CreatedAt time.Time `gorm:"type:DATETIME;index"`
	UpdatedAt time.Time `gorm:"type:DATETIME"`
}

type UIDModel struct {
	UID
	Time
}

type IIDModel struct {
	IID
	Time
}

type CipherText string

// TODO replace with secure key
var cipherKey = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

func (a *CipherText) Scan(v any) error {
	text, ok := v.([]byte)
	if !ok {
		return errors.New("invalid cipher text")
	}
	if len(text) == 0 {
		return nil
	}
	data, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(string(text))
	if err != nil {
		return err
	}
	key := cipherKey
	res, err := aesutil.Decrypt(data, key, key)
	if err != nil {
		return err
	}
	*a = CipherText(res)
	return nil
}

func (a CipherText) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "", nil
	}
	key := cipherKey
	data, err := aesutil.Encrypt([]byte(a), key, key)
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(data), nil
}

type Array []string

func (a *Array) Scan(v any) error {
	b := v.([]byte)
	return json.Unmarshal(b, a)
}

func (a Array) Value() (driver.Value, error) {
	return json.Marshal(a)
}
