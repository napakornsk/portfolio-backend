package main

import (
	"github.com/napakornsk/cv-backend/database"
	"github.com/napakornsk/cv-backend/entities"
	"gorm.io/gorm"
)

func main() {
	database.InitPostgresDB()
	db := database.PostgresDb
	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&entities.CV{},
		&entities.Education{},
		&entities.WorkDesc{},
		&entities.WorkExperience{},
		&entities.Skill{},
		&entities.Intro{},
	)
}
