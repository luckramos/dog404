package main

import (
	"dog404/middleware"
	"dog404/routes"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	routes.CreateAppRoutes(router)

	middlewaresGroup := middleware.MiddlewareGroup(
		middleware.CorsMiddleware,
		middleware.LoggingMiddleware,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: middlewaresGroup(router),
	}
	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
