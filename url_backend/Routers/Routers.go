package routers

import (
	"net/http"
    "io"
	// "alexshelto/url_shorten_service/Controllers"
)

func GetPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Get page!\n")
}

func ShortenUrl(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Get page!\n")
}


func SetUpRouters() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/api/v1/p", ShortenUrl)
    mux.HandleFunc("/p", GetPage)

    return mux
}




