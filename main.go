package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"alexshelto/url_shorten_service/database"
	"alexshelto/url_shorten_service/handlers"
	"alexshelto/url_shorten_service/services"
	"alexshelto/url_shorten_service/routes"
)


func main() {
    db := database.InitializeDatabase("data.db")

    linkService := services.NewLinkService(db)
    linkHandler := handlers.NewLinkHandler(linkService)

    // Create new router
    router := gin.Default()
    router.Use(cors.Default())

    // Apply Endpoints 
    routes.SetupRoutes(router, linkHandler)

    // Listen and serve
    router.Run(":8080")
}

