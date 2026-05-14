package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk koneksi database JapaneseClub
var DB *gorm.DB

// DataSiswa adalah model GORM untuk tabel data_siswas
type DataSiswa struct {
	// gorm.Model otomatis memberikan kolom ID (id), CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Nama         string    `json:"nama" gorm:"type:varchar(100);not null"`
	NoHP         string    `json:"no_hp" gorm:"type:varchar(15);not null"`
	TanggalLahir time.Time `json:"tanggal_lahir" gorm:"type:date;not null"` 
	Kelas        string    `json:"kelas" gorm:"type:varchar(30);not null"`         
	Divisi       string    `json:"divisi" gorm:"type:varchar(50);not null"`
}

// ConnectDatabase berfungsi untuk membuka koneksi ke PostgreSQL dan melakukan migrasi tabel
func ConnectDatabase() {
	// Konfigurasi DSN (Data Source Name) untuk database JapaneseClub
	host := "localhost"
	user := "postgres"
	password := "ghulam159_"       // Ganti dengan password PostgreSQL kamu
	dbname := "JapaneseClub"    // Nama database sesuai request kamu
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", 
		host, user, password, dbname, port)
	
	// Membuka koneksi menggunakan GORM
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terkoneksi ke database PostgreSQL: ", err)
	}

	// AutoMigrate akan otomatis membuat tabel 'data_siswas' di database JapaneseClub
	err = database.AutoMigrate(&DataSiswa{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database: ", err)
	}

	fmt.Println("Berhasil terhubung ke database JapaneseClub dan migrasi tabel DataSiswa selesai!")
	
	// Masukkan koneksi ke variabel global DB
	DB = database
}