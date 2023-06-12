package handlers

import (
	"errors"
    "log"
	"net/http"
    "strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (lh *LinkHandler) GetLinkById(context *gin.Context) {
	linkId := context.Param("id")

    // Cast String As uint 
    uintID, err := strconv.ParseUint(linkId, 10, 64)
    if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
        return
    }

    log.Println("No error converting to uint")

    log.Println("Successfully conveted to uint")
	link, err := lh.LinkService.GetLinkById(uint(uintID))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "404 Resource not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch link"})
		}
		return
	}

    log.Println(http.StatusOK)

    context.JSON(http.StatusOK, link)
}

