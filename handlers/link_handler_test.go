package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"

	"alexshelto/url_shorten_service/repositories"
	"alexshelto/url_shorten_service/routes"
	"alexshelto/url_shorten_service/services"
	"alexshelto/url_shorten_service/testutils"
)

func setupTest() (*gorm.DB, *LinkHandler) {
	db := testutils.SetupTestDatabase()

	linkRepository := repositories.NewLinkRepository(db)
	linkService := services.NewLinkService(linkRepository)
	linkHandler := NewLinkHandler(linkService)

	return db, linkHandler
}


func TestLinkHandlerCreateLinkRequest_Fails_NoBody(t *testing.T) {
	db, linkHandler := setupTest()
	defer testutils.TeardownTestDatabase(db)

	router := gin.New()
	routes.SetupRoutes(router, linkHandler)

	req, err := http.NewRequest("POST", "/l", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Capture response
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, recorder.Code, http.StatusBadRequest)
}

func TestLinkHandlerCreateLinkRequest_Success(t *testing.T) {
	db, linkHandler := setupTest()
	defer testutils.TeardownTestDatabase(db)

	router := gin.New()
	routes.SetupRoutes(router, linkHandler)

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

	// Capture response
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, recorder.Code, http.StatusCreated)
}


func TestLinkHandlerGetLinkById_Success(t *testing.T) {
	db, linkHandler := setupTest()
	defer testutils.TeardownTestDatabase(db)

	router := gin.New()
	routes.SetupRoutes(router, linkHandler)

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
	assert.Equal(t, recorder.Code, http.StatusCreated)


    // Parse ID from created Link and get by ID

	var response struct {
		ID uint `json:"id"`
	}

	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to parse response body JSON: %v", err)
		return
	}

	url := fmt.Sprintf("/link/%d", response.ID)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

    recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
    assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestLinkHandlerGetLinkById_Fails(t *testing.T) {
	db, linkHandler := setupTest()
	defer testutils.TeardownTestDatabase(db)

	router := gin.New()
	routes.SetupRoutes(router, linkHandler)

	req, err := http.NewRequest("GET", "/link/12", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Capture response
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, recorder.Code, http.StatusNotFound)
}



func TestLinkHandlerGetLinkById_FailsBadId(t *testing.T) {
	db, linkHandler := setupTest()
	defer testutils.TeardownTestDatabase(db)

	router := gin.Default()
    gin.SetMode(gin.ReleaseMode)

	routes.SetupRoutes(router, linkHandler)

	req, err := http.NewRequest("GET", "/link/12a1afasf", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Capture response
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
