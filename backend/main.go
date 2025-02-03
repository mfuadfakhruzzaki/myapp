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
	// Inisialisasi koneksi database dan auto migration
	utils.InitDB()

	router := mux.NewRouter()

	// Endpoint untuk register dan login tidak memerlukan autentikasi
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Endpoint CRUD untuk "items", dilindungi oleh middleware JWT
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTAuth)
	protected.HandleFunc("/items", handlers.GetItemsHandler).Methods("GET")
	protected.HandleFunc("/items", handlers.CreateItemHandler).Methods("POST")
	protected.HandleFunc("/items/{id}", handlers.UpdateItemHandler).Methods("PUT")
	protected.HandleFunc("/items/{id}", handlers.DeleteItemHandler).Methods("DELETE")

	fmt.Println("Backend berjalan di port 8081...")
    log.Fatal(http.ListenAndServe(":8081", router))

}
