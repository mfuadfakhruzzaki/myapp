package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// jwtKey adalah kunci rahasia untuk menandatangani token JWT.
// Pastikan nilai ini sama dengan yang Anda gunakan di handler.
var jwtKey = []byte("secret_key")

// contextKey adalah tipe kustom untuk context agar tidak terjadi bentrok.
type contextKey string

// ContextUserKey digunakan untuk menyimpan username atau informasi user di context.
const ContextUserKey = contextKey("user")

// JWTAuth adalah middleware untuk memeriksa token JWT yang dikirim pada header Authorization.
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Dapatkan header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header tidak ditemukan", http.StatusUnauthorized)
			return
		}

		// Header harus dalam format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Format Authorization header tidak valid", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		// Buat objek klaim untuk menampung data token
		claims := &jwt.StandardClaims{}

		// Parse dan verifikasi token
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token tidak valid", http.StatusUnauthorized)
			return
		}

		// Simpan informasi user (misalnya, username) ke dalam context
		ctx := context.WithValue(r.Context(), ContextUserKey, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
