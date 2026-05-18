package main // <-- WAJIB 'main' agar bisa dieksekusi lewat perintah 'go run'

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// Memanggil package lokal sesuai dengan nama module di go.mod kamu
	"backend-1/config"   // Memanggil folder config (cors)
	"backend-1/database" // Memanggil folder database
	"backend-1/routes"   // Memanggil folder routes
)

func main() {
	fmt.Println("=== Memulai Aplikasi JapaneseClub Backend ===")

	// 1. Nyalakan Koneksi Database (Membaca file database/database.go)
	database.ConnectDatabase()

	// 2. Buat Instance Router dari Gin Framework
	router := gin.Default()

	// 3. Pasang Satpam CORS Global (Membaca file config/cors.go)
	router.Use(config.SetupCORS())

	// 4. Daftarkan Alamat API (Membaca file routes/routes.go)
	routes.SetupRoutes(router)

	// 5. Jalankan Server di Port 8080
	fmt.Println("\n[SUCCESS] Server berjalan mulus di http://localhost:8080")
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}