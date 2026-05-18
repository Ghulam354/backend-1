package middleware

import (
	"backend-1/controllers"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil header Authorization dari request
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Akses ditolak, silakan login terlebih dahulu"})
			c.Abort() // Menghentikan eksekusi handler rute selanjutnya
			return
		}

		// 2. Format token haruslah "Bearer <string-token>"
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// 3. Validasi dan uraikan token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metode signing tidak valid")
			}
			return controllers.JwtSecret, nil
		})

		// 4. Jika token kadaluwarsa, rusak, atau palsu
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Sesi Anda habis atau token tidak valid, silakan login ulang"})
			c.Abort()
			return
		}

		// Jika lolos seleksi token, lanjutkan ke controller tujuan
		c.Next()
	}
}