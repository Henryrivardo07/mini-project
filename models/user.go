package models

// User mendefinisikan struktur data untuk pengguna
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
