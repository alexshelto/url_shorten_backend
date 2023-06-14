package repositories

import (
    "errors"
	"testing"

	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"

	"alexshelto/url_shorten_service/testutils"
	"alexshelto/url_shorten_service/models"
)


func setupTest() (*gorm.DB, *LinkRepository) {
    db := testutils.SetupTestDatabase()

    linkRepo := NewLinkRepository(db)

    return db, linkRepo;
}

func generateLink() models.Link {
    return models.Link {
        OriginalUrl: "https://www.github.com/alexshelto",
        ShortenedUrl: "123",
        VisitCount: 0,
    }
}


/// Defining Tests
func TestLinkRepositoryCreateLink_Success(t *testing.T) {
    db, linkRepo := setupTest()
    defer testutils.TeardownTestDatabase(db)
    
    link := generateLink()
    createdLink, err := linkRepo.Create(link)

    assert.NoError(t, err)
    assert.NotNil(t, createdLink)
    assert.NotEmpty(t, createdLink.ID)
    assert.Equal(t, createdLink.OriginalUrl, link.OriginalUrl)
}


func TestLinkRepositoryGetLinkById_Success(t *testing.T) {
    db, linkRepo := setupTest()
    defer testutils.TeardownTestDatabase(db)
    
    link := generateLink()

    createdLink, err := linkRepo.Create(link)

    assert.NoError(t, err)
    assert.NotNil(t, createdLink)
    assert.NotEmpty(t, createdLink.ID)
    assert.Equal(t, createdLink.OriginalUrl, link.OriginalUrl)

    var existingId uint = createdLink.ID

    requestedResource, err := linkRepo.GetById(existingId)

    assert.NoError(t, err)
    assert.NotNil(t, requestedResource)
    assert.NotEmpty(t, requestedResource.ID)
    assert.Equal(t, requestedResource.OriginalUrl, requestedResource.OriginalUrl)
    assert.Equal(t, requestedResource.ID, createdLink.ID)
    assert.Equal(t, requestedResource.OriginalUrl, createdLink.OriginalUrl)
}


func TestLinkRepositoryGetLinkById_Fails(t *testing.T) {
    db, linkRepo := setupTest()
    defer testutils.TeardownTestDatabase(db)
    
    var nonExistantId uint = 12
    _, err := linkRepo.GetById(nonExistantId)

    assert.Error(t, err)
    assert.Equal(t, errors.Is(err, gorm.ErrRecordNotFound), true)
}

func TestLinkRepositoryGetLinkByShortenedUrl_Success(t *testing.T) {
    db, linkRepo := setupTest()
    defer testutils.TeardownTestDatabase(db)
    
    link := generateLink()

    createdLink, err := linkRepo.Create(link)

    assert.NoError(t, err)
    assert.NotNil(t, createdLink)
    assert.NotEmpty(t, createdLink.ID)
    assert.Equal(t, createdLink.OriginalUrl, link.OriginalUrl)

    requestedResource, err := linkRepo.GetByShortenedUrl(createdLink.ShortenedUrl)

    assert.NoError(t, err)
    assert.NotNil(t, requestedResource)
    assert.NotEmpty(t, requestedResource.ID)
    assert.Equal(t, requestedResource.ID, createdLink.ID)
    assert.Equal(t, requestedResource.OriginalUrl, createdLink.OriginalUrl)
}


func TestLinkRepositoryGetLinkByShortenedUrl_Fails(t *testing.T) {
    db, linkRepo := setupTest()
    defer testutils.TeardownTestDatabase(db)
    
    _, err := linkRepo.GetByShortenedUrl("this-doesnt-exist")

    assert.Error(t, err)
    assert.Equal(t, errors.Is(err, gorm.ErrRecordNotFound), true)
}

