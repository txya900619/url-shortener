package kgs

import (
	"github.com/txya900619/url-shortener/internal/kgs/schema"
	"gorm.io/gorm"
)

const (
	source = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func insertUnusedKeys(prefix string, len int, db *gorm.DB) error {
	if len == 1 {
		keys := make([]schema.UnusedKey, 62)
		for i, c := range source {
			keys[i].Key = prefix + string(c)
		}
		return db.CreateInBatches(keys, 62).Error
	}
	for _, c := range source {
		err := insertUnusedKeys(prefix+string(c), len-1, db)
		if err != nil {
			return err
		}
	}

	return nil
}

func InsertUnusedKeys(db *gorm.DB) error {
	if err := db.First(&schema.UnusedKey{}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}

		err = db.Transaction(func(tx *gorm.DB) error {
			if err := insertUnusedKeys("", 6, tx); err != nil {
				tx.Rollback()
				return err
			}
			return tx.Commit().Error

		})
		return err
	}

	return nil
}
