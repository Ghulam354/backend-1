package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// 1. Definisikan handler
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Koneksi Backend Berhasil!"})
	})

	// 2. Tambahkan Println SEBELUM ListenAndServe
	fmt.Println("Server Go sedang berjalan di http://localhost:8080")
	fmt.Println("Tekan Ctrl+C untuk menghentikan server")

	// 3. Jalankan server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}