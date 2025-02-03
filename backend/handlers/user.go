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
// Disarankan untuk mengambil nilainya dari environment variable untuk keamanan.
var jwtKey = []byte("secret_key")

// RegisterHandler menangani pendaftaran user baru.
// Endpoint: POST /register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Decode request body ke struct User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Input tidak valid", http.StatusBadRequest)
		return
	}

	// Hash password sebelum menyimpan (pastikan password tidak disimpan dalam bentuk plaintext)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error pada proses hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Simpan user menggunakan GORM
	if result := utils.DB.Create(&user); result.Error != nil {
		http.Error(w, "Error menyimpan user: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LoginHandler menangani proses autentikasi user.
// Endpoint: POST /login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	// Decode JSON request body ke struct User (hanya memerlukan username dan password)
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Input tidak valid", http.StatusBadRequest)
		return
	}

	var user models.User
	// Cari user berdasarkan username menggunakan GORM
	if result := utils.DB.Where("username = ?", credentials.Username).First(&user); result.Error != nil {
		http.Error(w, "User tidak ditemukan", http.StatusUnauthorized)
		return
	}

	// Bandingkan password yang diinput dengan password yang disimpan (yang sudah di-hash)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Password salah", http.StatusUnauthorized)
		return
	}

	// Buat token JWT dengan masa berlaku 24 jam
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   user.Username,
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error membuat token", http.StatusInternalServerError)
		return
	}

	// Kirim token sebagai response dalam format JSON
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
