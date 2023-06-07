package main

import (
	"alexshelto/url_shorten_service/Models"
	"alexshelto/url_shorten_service/Routers"
)


func main() {
    router := routers.SetUpRouters()
    models.ConnectDatabase()

    router.Run(":8080")
}

