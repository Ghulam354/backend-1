package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	
	// Import modul internal kamu (ganti 'backend-1' sesuai nama go.mod-mu)
	"backend-1/config"   // Memanggil File 1 (CORS)
	"backend-1/database" // Memanggil File 3 (Database GORM)
	"backend-1/routes"   // Memanggil File 2 (Routing)
)

func main() {
	fmt.Println("Memulai inisialisasi server...")

	// 1. Nyalakan Koneksi Database (Membaca file database/db.go)
	database.ConnectDatabase()

	// 2. Buat Instance Router dari Gin Framework
	router := gin.Default()

	// 3. Pasang Satpam CORS Global (Membaca file config/cors.go)
	router.Use(config.SetupCORS())

	// 4. Daftarkan Alamat API (Membaca file routes/routes.go)
	routes.SetupRoutes(router)

	// 5. Jalankan Server di Port 8080
	fmt.Println("Server berjalan mulus di http://localhost:8080")
	router.Run(":8080")
}