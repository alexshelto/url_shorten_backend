package routes

import (
    "github.com/gin-gonic/gin"
)


type LinkHandlerInterface interface {
    CreateLink(context *gin.Context)
    GetLinkById(context *gin.Context) 
    RedirectToLink(context *gin.Context)
}


func SetupRoutes(router *gin.Engine, linkHandler LinkHandlerInterface) {
    router.POST("/l", linkHandler.CreateLink)         // Create Link
    router.GET("/link/:id", linkHandler.GetLinkById)
    router.GET("/l/:id", linkHandler.RedirectToLink)         // Visit Link (redirect)

}
