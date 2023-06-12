package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"alexshelto/url_shorten_service/models"
	"alexshelto/url_shorten_service/services"
)

type LinkHandler struct {
    LinkService *services.LinkService
}

func NewLinkHandler(service *services.LinkService) *LinkHandler {
    return &LinkHandler {
        LinkService: service,
    }
}


func (lh *LinkHandler) GetLink(context *gin.Context) {
    linkId := context.Param("id")

    link, err := lh.LinkService.GetLinkById(linkId)

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
    var newLink models.Link

    if err := context.ShouldBindJSON(&newLink); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    err := lh.LinkService.CreateLink(newLink)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create link"})
        return
    }

    context.JSON(http.StatusCreated, newLink)
}

func (lh *LinkHandler) Hello(context *gin.Context) {
    linkId := context.Param("id")
    context.JSON(http.StatusOK, linkId)
}

