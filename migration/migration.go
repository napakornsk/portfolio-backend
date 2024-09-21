package main

import (
	"log"

	"github.com/napakornsk/cv-backend/config"
	"github.com/napakornsk/cv-backend/db"
	"gorm.io/gorm"
)

// Run starts the migration process.
func main() {
	cfg := config.LoadConfig()

	store := db.NewStore(cfg.DB)
	defer store.Close() // Ensure the connection is closed after migrations

	migrate(store.DB)
}

// migrate performs the database migrations.
func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
	// &models.User{}, // Add more models as needed
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully!")
}
