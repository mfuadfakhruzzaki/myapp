package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Pastikan menyimpan hash password, bukan password plaintext!
}
