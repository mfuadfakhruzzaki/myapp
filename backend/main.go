package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello from Golang Backend!"))
    })

    fmt.Println("Backend berjalan di port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Server error:", err)
    }
}