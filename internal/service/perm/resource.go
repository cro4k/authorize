package perm

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

type resources struct {
	simple   sync.Map
	patterns sync.Map
}

func (r *resources) init(db *gorm.DB) {}

func (r *resources) Find(url string) (string, error) {
	if v, ok := r.simple.Load(url); ok {
		return v.(string), nil
	}

	var matched string
	r.patterns.Range(func(key, value any) bool {
		if r.match(key, url) {
			matched = value.(string)
			return false
		}
		return true
	})
	if matched != "" {
		return matched, nil
	}

	return "", errors.New("not found")
}

func (r *resources) match(pattern any, url string) bool {

	//TODO
	return false
}
