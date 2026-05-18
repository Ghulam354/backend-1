package controllers

import (
	"backend-1/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// JwtSecret adalah kunci rahasia rahasia untuk tanda tangan token JWT.
// Jaga string ini agar tetap rahasia.
var JwtSecret = []byte("KUNCI_RAHASIA_JAPANESE_CLUB_SUPER_SECURE_123")

func LoginAdmin(c *gin.Context) {
	var input LoginInput

	// 1. Validasi struktur JSON request dari frontend
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Format data tidak sesuai"})
		return
	}

	// 2. Cari admin di database berdasarkan username
	var admin database.Admin
	if err := database.DB.Where("username = ?", input.Username).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
		return
	}

	// 3. Verifikasi Password (menggunakan teks mentah sesuai database.go kamu)
	if admin.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
		return
	}

	// 4. Buat Token JWT jika data cocok (Sesi aktif selama 24 jam)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": admin.ID,
		"username": admin.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat token keamanan"})
		return
	}

	// 5. Kirim token JWT ke frontend
	c.JSON(http.StatusOK, gin.H{
		"message": "Login sukses!",
		"token":   tokenString,
	})
}