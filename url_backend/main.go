package main

// TODO: Use net/http instead of fiber

import (
    "net/http"
    "alexshelto/url_shorten_service/Routers"
)


func main() {
    mux := routers.SetUpRouters()
    if err := http.ListenAndServe(":8000", mux); err != nil {
    panic(err)
  }
}

