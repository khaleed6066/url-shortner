package routes

import (
	"net/http"
	"url-shortner/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.GetAllShortUrl)
	http.HandleFunc("/shorten", handlers.ShortUrlHandler)
	http.HandleFunc("/redirect/", handlers.RedirectUrl)
}
