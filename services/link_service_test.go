package services

import (
    "errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"alexshelto/url_shorten_service/testutils"
	"alexshelto/url_shorten_service/models"
	"alexshelto/url_shorten_service/repositories"
)



func setupTest() (*gorm.DB, *LinkService) {
    db := testutils.SetupTestDatabase()

    linkRepo := repositories.NewLinkRepository(db)
    linkService := NewLinkService(linkRepo)

    return db, linkService 
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
    t.Run("GetLinkByShortenedUrl_Success", TestGetLinkByShortenedUrl_Success)
    t.Run("GetLinkByShortenedUrl_Fails", TestGetLinkByShortenedUrl_Fails)
}



/// Defining Tests

func TestCreateLink_Success(t *testing.T) {
    db, linkService := setupTest()
    defer testutils.TeardownTestDatabase(db)

    linkPayload := generateLink()

    createdLink, err := linkService.CreateLink(linkPayload)

    assert.NoError(t, err)
    assert.NotEmpty(t, createdLink.ID)
    assert.Equal(t, createdLink.OriginalUrl, linkPayload.OriginalUrl)
}


func TestGetLinkById_Success(t *testing.T) {
    db, linkService := setupTest()
    defer testutils.TeardownTestDatabase(db)

    linkPayload := generateLink()

    createdLink, err := linkService.CreateLink(linkPayload)
    assert.NoError(t, err)

    requestedResource, err := linkService.GetLinkById(createdLink.ID)
    assert.NoError(t, err)
    assert.Equal(t, requestedResource.OriginalUrl, createdLink.OriginalUrl)

}

func TestGetLinkById_Fails(t *testing.T) {
    db, linkService := setupTest()
    defer testutils.TeardownTestDatabase(db)

    var nonExistantId uint = 12;

    requestedResource, err := linkService.GetLinkById(nonExistantId)
    assert.Error(t, err)
    assert.Equal(t, errors.Is(err, gorm.ErrRecordNotFound), true)
    assert.Empty(t, requestedResource)
}


func TestGetLinkByShortenedUrl_Success(t *testing.T) {
    db, linkService := setupTest()
    defer testutils.TeardownTestDatabase(db)

    linkPayload := generateLink()

    createdLink, err := linkService.CreateLink(linkPayload)
    assert.NoError(t, err)

    requestedResource, err := linkService.GetLinkByShortenedUrl(createdLink.ShortenedUrl)
    assert.NoError(t, err)
    assert.NotNil(t, requestedResource)
    assert.Equal(t, requestedResource.OriginalUrl, createdLink.OriginalUrl)
    assert.Equal(t, requestedResource.ID, createdLink.ID)
}

func TestGetLinkByShortenedUrl_Fails(t *testing.T) {
    db, linkService := setupTest()
    defer testutils.TeardownTestDatabase(db)

    _, err := linkService.GetLinkByShortenedUrl("this-doesnt-exist")

    assert.Error(t, err)
    assert.Equal(t, errors.Is(err, gorm.ErrRecordNotFound), true)
}

