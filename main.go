package main


import (
    "alexshelto/url_shorten_service/Routers"
    "alexshelto/url_shorten_service/Models"
)


func main() {
    router := routers.SetUpRouters()
    models.ConnectDatabase()

    router.Run(":8080")
}

