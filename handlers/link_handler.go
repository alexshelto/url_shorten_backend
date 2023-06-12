package handlers

import (
	"errors"
    "log"
	"net/http"
    //"strconv"

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


func (lh *LinkHandler) RedirectToLink(context *gin.Context) {
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



func (lh *LinkHandler) CreateLink(context *gin.Context) {
    log.Println("\n\n\nInside of create link handler")


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

    log.Println("Status Created")
	context.JSON(http.StatusCreated, createdLink)
}


