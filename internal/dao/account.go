package dao

import (
	"github.com/cro4k/authorize/internal/model"
	"github.com/cro4k/authorize/utils/reg"
	"github.com/cro4k/common/crypto/hashutil"
	"gorm.io/gorm"
)

var Account = new(accountService)

type accountService struct{}

func (s *accountService) GetByID(accountID string) (*model.Account, error) {
	var acc = new(model.Account)
	return acc, id(accountID, acc)
}

func (s *accountService) Find(username string, selects ...string) (*model.Account, error) {
	var query = db
	if reg.Cellphone.MatchString(username) {
		query = query.Where("cellphone_hash = ?", hashutil.MD5s(username))
	} else if reg.Email.MatchString(username) {
		query = query.Where("email_hash = ?", hashutil.MD5s(username))
	} else {
		query = query.Where("username = ?", username)
	}
	for _, sel := range selects {
		query = query.Select(sel)
	}
	var acc = new(model.Account)
	err := query.First(acc).Error
	return acc, err
}

func (s *accountService) Register(acc *model.Account, profile *model.AccountProfile) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(acc).Error; err != nil {
			return err
		}
		profile.ID = acc.ID
		if err := tx.Create(profile).Error; err != nil {
			return err
		}
		return nil

	})
}

func (s *accountService) Profile(id string) (*model.AccountProfile, error) {
	var profile = new(model.AccountProfile)
	err := db.Where("id = ?", id).First(&profile).Error
	return profile, err
}
