package routes

import (
    "github.com/gin-gonic/gin"
)


type LinkHandlerInterface interface {
    GetLinkByShortenedUrl(context *gin.Context)
    CreateLink(context *gin.Context)
    GetLinkById(context *gin.Context) 
}


func SetupRoutes(router *gin.Engine, linkHandler LinkHandlerInterface) {
    router.POST("/l", linkHandler.CreateLink)         // Create Link
    router.GET("/link/:id", linkHandler.GetLinkById)
    router.GET("/l/:id", linkHandler.GetLinkByShortenedUrl)         // Visit Link (redirect)

}
