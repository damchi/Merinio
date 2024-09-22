package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"merinio/api"
	postgresgorm "merinio/pkg"
	"os"
	"strconv"
	"time"
)

func initDB() (*gorm.DB, error) {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	cfg := postgresgorm.Config{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     port,
		Schema:   os.Getenv("DB_NAME"),
	}

	var gormDb *gorm.DB
	var err error

	// Retry logic
	for i := 0; i < 5; i++ {
		gormDb, err = postgresgorm.InitPostgresGorm(cfg)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Printf("Failed to connect to database after retries: %v", err)
		return nil, err
	}
	return gormDb, nil
}

func main() {

	_, err := initDB()
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return
	}
	r := gin.New()
	r.Use(gin.Recovery())

	api.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Printf("Failed to run server: %v", err)
	}
}
