package main

import (
	"alexshelto/url_shorten_service/Models"
	"alexshelto/url_shorten_service/Routers"
	"os"

    "log"
	"github.com/joho/godotenv"
)


func main() {
    router := routers.SetUpRouters()
    models.ConnectDatabase()

    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Some error occured. Err: %s", err)
    }

    port := os.Getenv("port")  
    if port == "" {
        port = "8080"
    }

    router.Run(":"+port)
}

