package gorm

import (
	"fmt"

	"github.com/spf13/viper"
	gormotel "github.com/wei840222/gorm-otel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

func Open() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Use(gormotel.New(gormotel.WithLogResult(true), gormotel.WithSqlParameters(true)))
	if err != nil {
		return nil, err
	}

	err = db.Use(prometheus.New(prometheus.Config{}))
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

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Taipei", host, user, password, dbName, port)
}
