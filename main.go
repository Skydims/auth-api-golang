// main.go
package main

import (
	"log"
	"net/http"

	"auth-api/database"
	"auth-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB() // Inisialisasi database

	r := mux.NewRouter()
	routes.RegisterAuthRoutes(r)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
