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
    // Inisialisasi koneksi database (jika menggunakan database, misalnya PostgreSQL)
    utils.InitDB()

    router := mux.NewRouter()

    // Endpoint untuk register dan login (CRUD & autentikasi)
    router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
    router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

    // Endpoint CRUD untuk resource (misalnya, "items")
    router.Handle("/items", middleware.JWTAuth(http.HandlerFunc(handlers.GetItemsHandler))).Methods("GET")
    router.Handle("/items", middleware.JWTAuth(http.HandlerFunc(handlers.CreateItemHandler))).Methods("POST")
    router.Handle("/items/{id}", middleware.JWTAuth(http.HandlerFunc(handlers.UpdateItemHandler))).Methods("PUT")
    router.Handle("/items/{id}", middleware.JWTAuth(http.HandlerFunc(handlers.DeleteItemHandler))).Methods("DELETE")

    fmt.Println("Backend berjalan di port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal("Server error:", err)
    }
}
