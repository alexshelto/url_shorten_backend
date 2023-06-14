package routes

import (
    "github.com/gin-gonic/gin"
)


type LinkHandlerInterface interface {
    CreateLink(context *gin.Context)
    CreateLinkPage(context *gin.Context)
    CreateLinkFormHandler(context *gin.Context)

    GetLinkById(context *gin.Context) 
    RedirectToLink(context *gin.Context)
    GetAnalyticsByUrl(context *gin.Context)
}


// TODO: Really Drill in difference between "ID" and "shortened_url" or "encoded_url"
func SetupRoutes(router *gin.Engine, linkHandler LinkHandlerInterface) {

    // Crud Operations
    router.POST("/l", linkHandler.CreateLink)         // Create Link
    router.GET("/link/:id", linkHandler.GetLinkById)  //TODO: Change to ID 

    // Forms 
    // Dont want a collision
    router.GET("/l/create/form", linkHandler.CreateLinkPage)
    router.POST("/l/create/form", linkHandler.CreateLinkFormHandler)

    // Redirect and analytics routes
    router.GET("/l/:hash", linkHandler.RedirectToLink)         
    router.GET("/l/analytics/:hash", linkHandler.GetAnalyticsByUrl)         
}
