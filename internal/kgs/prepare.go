package kgs

import (
	"github.com/spf13/viper"
	"github.com/txya900619/url-shortener/internal/kgs/schema"
	"gorm.io/gorm"
)

const (
	source = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func insertUnusedKeys(prefix string, length int, db *gorm.DB) error {
	if length == 1 {
		keys := make([]schema.UnusedKey, 62)
		for i, c := range source {
			keys[i].Key = prefix + string(c)
		}
		return db.CreateInBatches(keys, 62).Error
	}
	for _, c := range source {
		err := insertUnusedKeys(prefix+string(c), length-1, db)
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

		keyLength := viper.GetInt("KEY_LENGTH")

		if err := insertUnusedKeys("", keyLength, db); err != nil {
			return err
		}

		return nil
	}

	return nil
}
