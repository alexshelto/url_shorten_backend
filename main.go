package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"alexshelto/url_shorten_service/handlers"
	"alexshelto/url_shorten_service/models"
	"alexshelto/url_shorten_service/services"
	"alexshelto/url_shorten_service/routes"
)


func main() {
    db := initializeDatabase("data.db")

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

func initializeDatabase(db_name string) *gorm.DB {
    db, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    })

	if err != nil {
		panic("Failed to connect to database")
	}

    // Auto-migrate 
    err = db.AutoMigrate(&models.Link{})
    if err != nil {
        panic("Failed to migrate database")
    }

	return db
}
