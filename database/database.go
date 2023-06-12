package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"alexshelto/url_shorten_service/models"
)

func InitializeDatabase(db_name string) *gorm.DB {
    db, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    })

	if err != nil {
		panic("Failed to connect to database")
	}

    // Auto-migrate 
    err = db.AutoMigrate(&models.Link{})
    if err != nil {
        panic("Failed to migrate database")
    }

	return db
}
