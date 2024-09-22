package postgres_gorm

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"strconv"
	"testing"
)

var testDb *gorm.DB

func TestMain(m *testing.M) {
	port, _ := strconv.Atoi(os.Getenv("DB_CONFIG_PORT"))
	fmt.Print(port)
	cfg := Config{
		Host:     "localhost",
		Username: "postgres",
		Password: "postgres",
		Port:     5432,
		Schema:   "merinio",
	}

	var err error
	testDb, err = InitPostgresGorm(cfg)
	if err != nil {
		os.Exit(1)
	}
	code := m.Run()
	sqlDb, _ := testDb.DB()
	sqlDb.Close()
	os.Exit(code)
}
func TestGetConnection(t *testing.T) {
	conn := GetConnection()
	if conn == nil {
		t.Fatalf("Expected non-nil connection, got nil")
	}
}

func TestCheckConnection(t *testing.T) {
	if !CheckConnection() {
		t.Fatal("Expected connection to be alive, but it is not")
	}
}
