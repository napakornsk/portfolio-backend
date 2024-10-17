package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDb *gorm.DB

func InitPostgresDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatalf("DB_HOST not set in .env")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatalf("DB_USER not set in .env")
	}

	password := os.Getenv("DB_PASSWORD")
	if user == "" {
		log.Fatalf("DB_PASSWORD not set in .env")
	}
	dbName := os.Getenv("DB_NAME")
	if user == "" {
		log.Fatalf("DB_NAME not set in .env")
	}
	port := os.Getenv("DB_PORT")
	if user == "" {
		log.Fatalf("DB_PORT not set in .env")
	}
	timezone := os.Getenv("TIMEZONE")
	if user == "" {
		log.Fatalf("TIMEZONE not set in .env")
	}
	sslMode := os.Getenv("SSL_MODE")
	if user == "" {
		log.Fatalf("SSL_MODE not set in .env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbName, port, sslMode, timezone,
	)
	fmt.Println("DSN: ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	PostgresDb = db
}
