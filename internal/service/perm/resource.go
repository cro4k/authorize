package perm

import (
	"errors"
	"strconv"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/cro4k/authorize/common/utils/pathutil"
	"github.com/cro4k/authorize/internal/model"
)

type resources struct {
	simple   sync.Map
	patterns sync.Map
}

func (r *resources) init(db *gorm.DB) {
	var apis []*model.PermApi
	db.Preload("Resource").Find(&apis)

	for _, api := range apis {
		r.simple.Store(api.Path, api.ResourceID)
		if strings.Contains(api.Path, "(") || strings.Contains(api.Path, "{") {
			p, err := pathutil.NewURLPattern(api.Path)
			if err == nil {
				r.patterns.Store(p, api.ResourceID)
			} else {
				logrus.Errorf("%d %v", api.ID, err)
			}
		}
	}
}

func (r *resources) Find(url string) (string, error) {
	if v, ok := r.simple.Load(url); ok {
		return strconv.Itoa(int(v.(uint64))), nil
	}

	var matched string
	r.patterns.Range(func(key, value any) bool {
		if r.match(key, url) {
			matched = strconv.Itoa(int(value.(uint64)))
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
	p, ok := pattern.(*pathutil.URLPattern)
	if !ok {
		return false
	}
	return p.Match(url)
}
