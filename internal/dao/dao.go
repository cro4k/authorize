package dao

import (
	"gorm.io/gorm"
)

func list(v interface{}, query *gorm.DB, page, size int, order ...interface{}) (int64, error) {
	var total int64
	query.Count(&total)
	for _, o := range order {
		query = query.Order(o)
	}
	query = query.Offset(size * (page - 1)).Limit(size)
	err := query.Find(v).Error
	return total, err
}

func id(db *gorm.DB, v interface{}, id interface{}, selects ...interface{}) error {
	query := db.Where("id = ?", id)
	for _, sel := range selects {
		query = query.Select(sel)
	}
	return query.First(v).Error
}
