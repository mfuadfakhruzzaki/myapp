package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mfuadfakhruzzaki/myapp-backend/models"
	"github.com/mfuadfakhruzzaki/myapp-backend/utils"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// jwtKey adalah kunci rahasia untuk menandatangani token JWT.
// Pastikan untuk menggunakan kunci yang aman di lingkungan produksi.
var jwtKey = []byte("secret_key") 

// RegisterHandler menangani pendaftaran user baru.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Decode JSON request body ke struct User.
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash password sebelum disimpan.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Simpan user ke database.
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err = utils.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LoginHandler menangani autentikasi user dan menghasilkan token JWT.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	// Decode kredensial login dari request.
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var storedUser models.User
	// Ambil data user dari database berdasarkan username.
	query := "SELECT id, username, password FROM users WHERE username=$1"
	err := utils.DB.QueryRow(query, creds.Username).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Bandingkan password yang diberikan dengan hash yang tersimpan.
	if err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Buat token JWT yang berlaku selama 24 jam.
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   storedUser.Username,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
