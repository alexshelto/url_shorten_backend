package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"alexshelto/url_shorten_service/database"
	"alexshelto/url_shorten_service/handlers"
	"alexshelto/url_shorten_service/services"
	"alexshelto/url_shorten_service/repositories"
	"alexshelto/url_shorten_service/routes"
)


func main() {
    db := database.InitializeDatabase("data.db")

    linkRepository := repositories.NewLinkRepository(db)
    linkService := services.NewLinkService(linkRepository)
    linkHandler := handlers.NewLinkHandler(linkService)

    // Create new router
    router := gin.Default()
    router.LoadHTMLGlob("static/*html")
    router.Use(cors.Default())

    // Apply Endpoints 
    routes.SetupRoutes(router, linkHandler)

    // Listen and serve
    router.Run(":8080")
}

