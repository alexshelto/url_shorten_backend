package services

import (
    "errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"alexshelto/url_shorten_service/database"
	"alexshelto/url_shorten_service/models"
	"alexshelto/url_shorten_service/repositories"
)


var TestDatabaseName = "test_database.db"


func setupTestDatabase() *gorm.DB {
    return database.InitializeDatabase(TestDatabaseName)
}

func setupTest() (*gorm.DB, *LinkService) {
    db := setupTestDatabase()
    linkRepo := repositories.NewLinkRepository(db)
    linkService := NewLinkService(linkRepo)

    return db, linkService 
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


func generateLink() models.Link {
    return models.Link {
        OriginalUrl: "https://www.github.com/alexshelto",
        ShortenedUrl: "123",
        VisitCount: 0,
    }
}

// Running Tests 
func TestLinkService(t *testing.T) {
    t.Run("CreateLink_Success", TestCreateLink_Success)
    t.Run("GetLinkById_Success", TestGetLinkById_Success)
    t.Run("GetLinkById_Fails", TestGetLinkById_Fails)
}



/// Defining Tests

func TestCreateLink_Success(t *testing.T) {
    db, linkService := setupTest()
    defer teardownTestDatabase(db)

    linkPayload := generateLink()

    createdLink, err := linkService.CreateLink(linkPayload)

    assert.NoError(t, err)
    assert.NotEmpty(t, createdLink.ID)
    assert.Equal(t, createdLink.OriginalUrl, linkPayload.OriginalUrl)
}


func TestGetLinkById_Success(t *testing.T) {
    db, linkService := setupTest()
    defer teardownTestDatabase(db)

    linkPayload := generateLink()

    createdLink, err := linkService.CreateLink(linkPayload)
    assert.NoError(t, err)

    requestedResource, err := linkService.GetLinkById(createdLink.ID)
    assert.NoError(t, err)
    assert.Equal(t, requestedResource.OriginalUrl, createdLink.OriginalUrl)

}

func TestGetLinkById_Fails(t *testing.T) {
    db, linkService := setupTest()
    defer teardownTestDatabase(db)

    var nonExistantId uint = 12;

    requestedResource, err := linkService.GetLinkById(nonExistantId)
    assert.Error(t, err)
    assert.Equal(t, errors.Is(err, gorm.ErrRecordNotFound), true)
    assert.Empty(t, requestedResource)
}


