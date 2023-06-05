package main


import (
    "net/http"
    "github.com/gin-gonic/gin"
)




func retrieve_page_from_hash(c *gin.Context) {
    name := c.Param("hash")
    c.String(http.StatusOK, "Hello %s", name)
}


func ping(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}


func main() {
    router := gin.Default()

    router.GET("/ping", ping)
    router.GET("/p/:hash", retrieve_page_from_hash)

    router.Run() // listen and serve on 0.0.0.0:8080
}
