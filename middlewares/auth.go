package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mini-project/utils"
	"strings"
)

// JWTAuth adalah middleware untuk memverifikasi token JWT
func JWTAuth(c *gin.Context) {
	// Mengambil token dari header Authorization
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		utils.ErrorResponse(c, "Authorization token is required")
		c.Abort() // Hentikan eksekusi lebih lanjut
		return
	}

	// Menghapus kata "Bearer " jika ada
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Parsing token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Memeriksa apakah token menggunakan metode signing yang benar
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// Jika token tidak menggunakan metode signing yang benar
			utils.ErrorResponse(c, "Invalid signing method")
			c.Abort() // Menghentikan eksekusi
			return nil, nil
		}
		return []byte("your_secret_key"), nil // Ganti dengan secret key Anda
	})

	// Menangani error jika parsing gagal atau token tidak valid
	if err != nil {
		utils.ErrorResponse(c, "Invalid token")
		c.Abort()
		return
	}

	// Jika token valid, lanjutkan eksekusi
	if !token.Valid {
		utils.ErrorResponse(c, "Token is not valid")
		c.Abort()
		return
	}

	// Token valid, lanjutkan ke handler berikutnya
	c.Next()
}
