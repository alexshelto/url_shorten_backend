package main

import (
    "github.com/gin-gonic/gin"
    "alexshelto/url_shorten_service/server"
)


func main() {

    router := gin.Default()
    server := server.NewServer(router)

    server.Start(":8080")

}

