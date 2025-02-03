package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mfuadfakhruzzaki/myapp-backend/handlers"
	"github.com/mfuadfakhruzzaki/myapp-backend/middleware"
	"github.com/mfuadfakhruzzaki/myapp-backend/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.InitDB()

	router := mux.NewRouter()

	// Tambahkan middleware CORS global
	router.Use(middleware.CORS)

	// Endpoint publik
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST", "OPTIONS")

	// Endpoint protected
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTAuth)
	protected.HandleFunc("/items", handlers.GetItemsHandler).Methods("GET", "OPTIONS")
	protected.HandleFunc("/items", handlers.CreateItemHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/items/{id}", handlers.UpdateItemHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/items/{id}", handlers.DeleteItemHandler).Methods("DELETE", "OPTIONS")

	fmt.Println("Server berjalan di port 8081...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
