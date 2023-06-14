package handlers

import (
	"bytes"
    "encoding/json"
	"net/http"
	"net/http/httptest"

	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"alexshelto/url_shorten_service/models"
	"alexshelto/url_shorten_service/repositories"
	"alexshelto/url_shorten_service/routes"
	"alexshelto/url_shorten_service/services"
	"alexshelto/url_shorten_service/testutils"
)

var TestLinks = []models.Link{
	models.Link{
		OriginalUrl:  "https://www.github.com/alexshelto",
		ShortenedUrl: "123",
		VisitCount:   0,
	},
	models.Link{
		OriginalUrl:  "https://www.twitch.tv/smoke",
		ShortenedUrl: "abc",
		VisitCount:   1,
	},
	models.Link{
		OriginalUrl:  "https://www.google.com",
		ShortenedUrl: "XYZ",
		VisitCount:   12,
	},
}

func setupTest() (*gorm.DB, *gin.Engine) {
	db := testutils.SetupTestDatabase()
	testutils.SeedLinksTestDatabase(db, TestLinks)

	linkRepository := repositories.NewLinkRepository(db)
	linkService := services.NewLinkService(linkRepository)
	linkHandler := NewLinkHandler(linkService)

    router := gin.New()
    gin.SetMode(gin.ReleaseMode)

    router.LoadHTMLGlob("../static/*.html")
    routes.SetupRoutes(router, linkHandler)

	return db, router
}


func TestLinkHandler_CreateLinkRequest_Fails_NoBody(t *testing.T) {
	db, router := setupTest()
	defer testutils.TeardownTestDatabase(db)

	req, err := http.NewRequest("POST", "/l", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestLinkHandler_CreateLinkRequest_Success(t *testing.T) {
	db, router := setupTest()
	defer testutils.TeardownTestDatabase(db)

	requestBody := map[string]interface{}{
		"original_url": "https://www.github.com/alexshelto",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/l", bytes.NewReader(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}


func TestLinkHandler_GetLinkById_Success(t *testing.T) {
	db, router := setupTest()
	defer testutils.TeardownTestDatabase(db)

	recorder := httptest.NewRecorder()

    req, err := http.NewRequest("GET", "/l/id/2", nil)
	if err != nil {
		t.Fatal(err)
	}

    recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
    assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestLinkHandler_GetLinkById_Fails_404(t *testing.T) {
	db, router := setupTest()
	defer testutils.TeardownTestDatabase(db)

	req, err := http.NewRequest("GET", "/l/id/12", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}



func TestLinkHandler_GetLinkById_Fails_BadId(t *testing.T) {
	db, router := setupTest()
	defer testutils.TeardownTestDatabase(db)

	req, err := http.NewRequest("GET", "/l/id/12a1afasf", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}


func TestLinkHandler_GetLinkByShortened_Redirect_Success(t *testing.T) {
	db, router := setupTest()
	defer testutils.TeardownTestDatabase(db)

    req, err := http.NewRequest("GET", "/l/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

    recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
    assert.Equal(t, http.StatusPermanentRedirect, recorder.Code)
}

func TestLinkHandler_GetLinkByShortened_Redirect_Fails(t *testing.T) {
	db, router := setupTest()
	defer testutils.TeardownTestDatabase(db)

	req, err := http.NewRequest("GET", "/l/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}


// Analytics 

func TestLinkHandler_GetLinkByShortened_Analytics_Success(t *testing.T) {
	db, router := setupTest()
	defer testutils.TeardownTestDatabase(db)

	req, err := http.NewRequest("GET", "/l/analytics/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}


func TestLinkHandler_GetLinkByShortened_Analytics_Fails_404(t *testing.T) {
	db, router := setupTest()
	defer testutils.TeardownTestDatabase(db)

	req, err := http.NewRequest("GET", "/l/analytics/this-doesnt-exist", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}




func TestLinkHandler_GetCreateLinkForm_Success(t *testing.T) {
    db, router := setupTest()
    defer testutils.TeardownTestDatabase(db)

	req, err := http.NewRequest("GET", "/l/create/form", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestLinkHandler_PostCreateLinkFormRedirects_Success(t *testing.T) {
    db, router := setupTest()
    defer testutils.TeardownTestDatabase(db)

    form := url.Values{}
    form.Add("url", "http://www.github.com/alexshelto")

	req, err := http.NewRequest("POST", "/l/create/form", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)
    assert.Equal(t, http.StatusOK, recorder.Code)
}


func TestLinkHandler_PostCreateLinkForm_BadFormData_Fails(t *testing.T) {
    db, router := setupTest()
    defer testutils.TeardownTestDatabase(db)

    form := url.Values{}
    form.Add("invalid-form-field-123", "blah-blah-blah")

	req, err := http.NewRequest("POST", "/l/create/form", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}

    recorder := httptest.NewRecorder()

    router.ServeHTTP(recorder, req)
    assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

