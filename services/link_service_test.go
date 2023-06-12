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

/*
// Running Tests 
func TestLinkService(t *testing.T) {
    t.Run("CreateLink_Success-Service", TestCreateLink_Success)
    t.Run("GetLinkById_Success-Service", TestGetLinkById_Success)
    t.Run("GetLinkById_Fails-Service", TestGetLinkById_Fails)
    t.Run("GetLinkByShortenedUrl_Success-Service", TestGetLinkByShortenedUrl_Success)
    t.Run("GetLinkByShortenedUrl_Fails-Service", TestGetLinkByShortenedUrl_Fails)
}
*/



/// Defining Tests
func TestLinkServiceCreateLink_Success(t *testing.T) {
    db, linkService := setupTest()
    defer testutils.TeardownTestDatabase(db)

    linkPayload := generateLink()

    createdLink, err := linkService.CreateLink(linkPayload)

    assert.NoError(t, err)
    assert.NotEmpty(t, createdLink.ID)
    assert.Equal(t, createdLink.OriginalUrl, linkPayload.OriginalUrl)
}


func TestLinkServiceGetLinkById_Success(t *testing.T) {
    db, linkService := setupTest()
    defer testutils.TeardownTestDatabase(db)

    linkPayload := generateLink()

    createdLink, err := linkService.CreateLink(linkPayload)
    assert.NoError(t, err)

    requestedResource, err := linkService.GetLinkById(createdLink.ID)
    assert.NoError(t, err)
    assert.Equal(t, requestedResource.OriginalUrl, createdLink.OriginalUrl)

}

func TestLinkServiceGetLinkById_Fails(t *testing.T) {
    db, linkService := setupTest()
    defer testutils.TeardownTestDatabase(db)

    var nonExistantId uint = 12;

    requestedResource, err := linkService.GetLinkById(nonExistantId)
    assert.Error(t, err)
    assert.Equal(t, errors.Is(err, gorm.ErrRecordNotFound), true)
    assert.Empty(t, requestedResource)
}


func TestLinkServiceGetLinkByShortenedUrl_Success(t *testing.T) {
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

func TestLinkServiceGetLinkByShortenedUrl_Fails(t *testing.T) {
    db, linkService := setupTest()
    defer testutils.TeardownTestDatabase(db)

    _, err := linkService.GetLinkByShortenedUrl("this-doesnt-exist")

    assert.Error(t, err)
    assert.Equal(t, errors.Is(err, gorm.ErrRecordNotFound), true)
}

