package db

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file :", err)
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN not set")
	}

	var error error
	DB, error = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}
	log.Println("DB connection established")
}
