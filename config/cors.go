package config

import (
	"github.com/gin-gonic/gin"
)

// SetupCORS mengatur kebijakan hak akses (CORS) agar frontend bisa mengakses backend
func SetupCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Mengizinkan asal request (Origin)
		// Menggunakan "*" berarti mengizinkan semua domain 
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		
		// 2. Mengizinkan pengiriman kredensial seperti cookie atau token auth
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		
		// 3. Menentukan HTTP Method apa saja yang diperbolehkan untuk memanipulasi API
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		
		// 4. Menentukan Header apa saja yang boleh dikirimkan oleh frontend
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		// 5. Menangani Preflight Request (Method OPTIONS)
		// Browser otomatis mengirimkan method OPTIONS sebelum mengirim request asli (POST/PUT)
		if c.Request.Method == "OPTIONS" {
			// Jika methodnya OPTIONS, langsung jawab dengan status 204 (No Content) dan hentikan proses di sini
			c.AbortWithStatus(204)
			return
		}

		// Melanjutkan request ke proses berikutnya (routing/controller)
		c.Next()
	}
}