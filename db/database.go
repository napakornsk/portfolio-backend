package db

import (
	"log"

	"github.com/napakornsk/cv-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

// NewStore creates a new Store instance with a GORM database connection.
func NewStore(config config.DBConfig) *Store {
	dsn := buildDSN(config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	return &Store{DB: db}
}

// buildDSN constructs a Data Source Name (DSN) for connecting to the database.
func buildDSN(config config.DBConfig) string {
	return "host=" + config.Host +
		" port=" + config.Port +
		" user=" + config.Username +
		" password=" + config.Password +
		" dbname=" + config.DbName +
		" sslmode=disable"
}

// Close closes the database connection.
func (store *Store) Close() {
	db, err := store.DB.DB()
	if err != nil {
		log.Fatal("failed to get generic database:", err)
	}
	if err := db.Close(); err != nil {
		log.Fatalf("failed to close the database: %v", err)
	}
}
