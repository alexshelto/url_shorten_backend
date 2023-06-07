package routers

import (
    "github.com/gin-gonic/gin"

	"alexshelto/url_shorten_service/Controllers"
)


func SetUpRouters() *gin.Engine {
    router := gin.Default()
    router.LoadHTMLGlob("static/*.html")

    // version 1
	apiV1 := router.Group("api/v1")
    apiV1.POST("/p", Controller.CreateHashedPageV1)
    router.GET("/p/:hash", Controller.GetPageFromHash)

    return router
}





