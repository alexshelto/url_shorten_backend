package handlers

import (
	"errors"
    "log"
	"net/http"
    "strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"alexshelto/url_shorten_service/models"
	"alexshelto/url_shorten_service/services"
)

type LinkHandler struct {
	LinkService *services.LinkService
}

func NewLinkHandler(service *services.LinkService) *LinkHandler {
	return &LinkHandler{
		LinkService: service,
	}
}

func (lh *LinkHandler) GetLinkByShortenedUrl(context *gin.Context) {
	linkId := context.Param("id")

	link, err := lh.LinkService.GetLinkByShortenedUrl(linkId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "404 Resource not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch link"})
		}
		return
	}

	context.Redirect(http.StatusPermanentRedirect, link.OriginalUrl)
}

func (lh *LinkHandler) GetLinkById(context *gin.Context) {
	linkId := context.Param("id")

    // Cast String As uint 
    uintID, err := strconv.ParseUint(linkId, 10, 64)
    if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
        return
    }

	link, err := lh.LinkService.GetLinkById(uint(uintID))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "404 Resource not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch link"})
		}
		return
	}

    log.Println("INSIDE OF get link by shortenedURL, ", link)
    context.JSON(http.StatusNotFound, link)
}




func (lh *LinkHandler) CreateLink(context *gin.Context) {
    var linkRequest models.CreateLinkRequest

	if err := context.ShouldBindJSON(&linkRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

    // Translate the Request into a GOrm link type
    link := models.Link {
        OriginalUrl: linkRequest.OriginalUrl,
    }

	createdLink, err := lh.LinkService.CreateLink(link)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create link"})
		return
	}

	context.JSON(http.StatusCreated, createdLink)
}


