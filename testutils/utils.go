package testutils


import (
	"log"
	"os"

	"gorm.io/gorm"

	"alexshelto/url_shorten_service/database"
	"alexshelto/url_shorten_service/models"
)


var TestDatabaseName = "test_database.db"

func SetupTestDatabase() *gorm.DB {
    db := database.InitializeDatabase(TestDatabaseName)
    return db
}

func SeedLinksTestDatabase(db *gorm.DB,links []models.Link) {
    err := db.Create(links).Error

    if err != nil {
        panic("Failed to seed test database, exiting")
    }
}

func TeardownTestDatabase(db *gorm.DB) {
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal("failed to get db connection", err)
    }
    sqlDB.Close()

    // remove db file 
    err = os.Remove(TestDatabaseName)
    if err != nil {
        log.Fatal("failed to remove test db file", err)
    }
}

