package orm

import (
	"github.com/txya900619/url-shortener/kgs/pkg/schema"

	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(getDsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&schema.UnusedKey{}, &schema.UsedKey{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDsn() string {
	host := viper.GetString("DB_HOST")
	port := viper.GetInt("DB_PORT")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Taiwan", host, user, password, dbName, port)
}
