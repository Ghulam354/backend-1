package routes

import (
	"backend-1/controllers"
	"backend-1/database"
	"backend-1/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// ==================== 1. ROUTE PUBLIK ====================
	// Jalur umum yang tidak dikunci JWT (Bisa diakses tanpa login)
	
	// Untuk kebutuhan indikator status koneksi di halaman Home frontend kamu
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Server JapaneseClub Aktif!"})
	})

	// Jalur untuk eksekusi login admin
	router.POST("/api/login", controllers.LoginAdmin)


	// ==================== 2. ROUTE TERPROTEKSI ====================
	// Jalur manajemen data siswa yang dipagari oleh AuthMiddleware (Wajib bawa token JWT)
	
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware()) // <--- 'Satpam' pengunci ditaruh di sini
	{
		// 1. Ambil Data Siswa (GET /api/siswa)
		api.GET("/siswa", func(c *gin.Context) {
			var siswa []database.DataSiswa
			if err := database.DB.Find(&siswa).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, siswa)
		})

		// 2. Tambah Data Siswa Baru (POST /api/siswa)
		api.POST("/siswa", func(c *gin.Context) {
			var input database.DataSiswa

			// Bind data JSON dari frontend ke struct DataSiswa
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Simpan data ke database PostgreSQL via GORM
			if err := database.DB.Create(&input).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Kirim respon balik ke frontend tanda sukses
			c.JSON(http.StatusOK, input)
		})
	}
}