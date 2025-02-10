package main

import (
	"HTTP-SERVER/database"
	"HTTP-SERVER/handlers"
	"HTTP-SERVER/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectDB()

	r := chi.NewRouter()

	r.Use(middleware.LogRequest)

	r.Post("/api/register", handlers.RegisterUser)

	r.Get("/api/users", handlers.GetUsers)

	port := ":8080"
	log.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
