package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"backend-1/database"
)

// InputSiswa adalah struktur data yang dikirim oleh frontend saat mendaftar
type InputSiswa struct {
	Nama         string `json:"nama" binding:"required"`
	NoHP         string `json:"no_hp" binding:"required"`
	TanggalLahir string `json:"tanggal_lahir" binding:"required"`
	Kelas        string `json:"kelas" binding:"required"`         
	Divisi       string `json:"divisi" binding:"required"`       
}

// SetupRoutes mendaftarkan semua endpoint API untuk JapaneseClub
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// 1. Endpoint untuk Mengambil Semua Data Siswa
		api.GET("/siswa", func(c *gin.Context) {
			var daftarSiswa []database.DataSiswa
			
			// Mengambil data dari database, diurutkan berdasarkan id paling baru
			if err := database.DB.Order("id desc").Find(&daftarSiswa).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari database"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Data siswa berhasil diambil",
				"data":    daftarSiswa,
			})
		})

		// 2. Endpoint untuk Menambah Data Siswa Baru
		api.POST("/siswa", func(c *gin.Context) {
			var input InputSiswa

			// Validasi input dari frontend, jika ada kolom yang kosong (required) akan error
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Semua kolom harus diisi dengan benar"})
				return
			}

			// Mengubah string "YYYY-MM-DD" dari frontend menjadi tipe time.Time agar bisa masuk ke PostgreSQL
			parsedDate, err := time.Parse("2006-01-02", input.TanggalLahir)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal lahir harus YYYY-MM-DD"})
				return
			}

			// Memasukkan data dari input frontend ke dalam struktur tabel database
			siswaBaru := database.DataSiswa{
				Nama:         input.Nama,
				NoHP:         input.NoHP,
				TanggalLahir: parsedDate,
				Kelas:        input.Kelas,
				Divisi:       input.Divisi,
			}

			// Menyimpan ke database menggunakan GORM
			if err := database.DB.Create(&siswaBaru).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data ke database"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Siswa berhasil didaftarkan ke JapaneseClub!",
				"data":    siswaBaru,
			})
		})
	}
}