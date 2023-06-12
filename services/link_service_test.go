package services

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"alexshelto/url_shorten_service/database"
	"alexshelto/url_shorten_service/models"
)


var TestDatabaseName = "test_database.db"


func setupTestDatabase() *gorm.DB {
    db := database.InitializeDatabase(TestDatabaseName)
    return db
}

func teardownTestDatabase(db *gorm.DB) {
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


// Running Tests 
func TestLinkService(t *testing.T) {
    t.Run("CreateLink_Success", TestCreateLink_Success)


}



// Defining Tests

func TestCreateLink_Success(t *testing.T) {
    db := setupTestDatabase()
    defer teardownTestDatabase(db)

    linkService := NewLinkService(db)

    linkPayload := models.Link{
        OriginalUrl: "https://www.github.com/alexshelto",
        ShortenedUrl: "foo",
        VisitCount: 0,
    }

    createdLink, err := linkService.CreateLink(linkPayload)

    assert.NoError(t, err)
    assert.NotEmpty(t, createdLink.ID)
    assert.Equal(t, createdLink.OriginalUrl, linkPayload.OriginalUrl)
}

