package main


import (
    "alexshelto/url_shorten_service/Routers"
)


func main() {
    router := routers.SetUpRouters()
    router.Run(":8000")
}

