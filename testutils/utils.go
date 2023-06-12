package testutils


import (
	"log"
	"os"

	"gorm.io/gorm"

	"alexshelto/url_shorten_service/database"
    /*
    "errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"alexshelto/url_shorten_service/models"
	"alexshelto/url_shorten_service/repositories"
    */
)



var TestDatabaseName = "test_database.db"


func SetupTestDatabase() *gorm.DB {
    return database.InitializeDatabase(TestDatabaseName)
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


