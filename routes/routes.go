package routes

import (
    "github.com/gin-gonic/gin"
    "alexshelto/url_shorten_service/handlers"
)


func SetupRoutes(router *gin.Engine, linkHandler *handlers.LinkHandler) {
    router.POST("/l", linkHandler.CreateLink)         // Create Link
    router.GET("/l/:id", linkHandler.GetLinkByShortenedUrl)         // Visit Link (redirect)
    router.GET("/l/stats/:id", linkHandler.GetLinkByShortenedUrl)   // Visit Link: stats with link
    router.GET("/hello/:id", linkHandler.Hello)       // Test Hello
}
