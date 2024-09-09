package routes

import (
	"dog404/handlers"
	"net/http"
)

func CreateAppRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /standard-404", handlers.StandardHandler)
	router.HandleFunc("POST /send-email", handlers.EmailHandler)

	router.Handle("GET /dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("../template/dist"))))
	router.Handle("GET /images/", http.StripPrefix("/images/", http.FileServer(http.Dir("../template/dist/images"))))
}
