package postgres_gorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	Host     string
	Port     int
	Schema   string
	Username string
	Password string
}

var gormDb *gorm.DB

func InitPostgresGorm(cfg Config) (*gorm.DB, error) {
	var err error

	databaseUrlFormat := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	databaseUrl := fmt.Sprintf(databaseUrlFormat, cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Schema)

	gormDb, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDb, nil
}

func GetConnection() *gorm.DB {
	return gormDb
}

func CheckConnection() bool {
	sqlDb, err := gormDb.DB()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	err = sqlDb.Ping()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
