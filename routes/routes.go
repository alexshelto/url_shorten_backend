package routes

import (
    "github.com/gin-gonic/gin"
)


type LinkHandlerInterface interface {
    CreateLink(context *gin.Context)
    GetLinkById(context *gin.Context) 
    RedirectToLink(context *gin.Context)
    GetAnalyticsByUrl(context *gin.Context)
}


// TODO: Really Drill in difference between "ID" and "shortened_url" or "encoded_url"
func SetupRoutes(router *gin.Engine, linkHandler LinkHandlerInterface) {
    router.POST("/l", linkHandler.CreateLink)         // Create Link
    router.GET("/link/:id", linkHandler.GetLinkById)

    router.GET("/l/:hash", linkHandler.RedirectToLink)         // Visit Link (redirect)
    router.GET("/l/analytics/:hash", linkHandler.GetAnalyticsByUrl)         // Visit Link (redirect)

}
