package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Table Admin
type Admin struct{
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	NoHP string `json:"no_hp" gorm:"type:varchar(15);not null"`
}

// Table DataSiswa
type DataSiswa struct {

	gorm.Model
	Nama         string    `json:"nama" gorm:"type:varchar(100);not null"`
	NoHP         string    `json:"no_hp" gorm:"type:varchar(15);not null"`
	TanggalLahir string `json:"tanggal_lahir" gorm:"type:date;not null"` 
	Kelas        string    `json:"kelas" gorm:"type:varchar(30);not null"`         
	Divisi       string    `json:"divisi" gorm:"type:varchar(50);not null"`
}

// Database Connection
func ConnectDatabase() {
	host := "localhost"
	user := "postgres"
	password := "rahasia"       
	dbname := "JapaneseClub"    
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", 
		host, user, password, dbname, port)
	 
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terkoneksi ke database PostgreSQL: ", err)
	}

	// AutoMigrate akan otomatis membuat tabel 'data_siswas' di database
	err = database.AutoMigrate(&DataSiswa{})
	err = database.AutoMigrate(&Admin{})

	if err != nil {
		log.Fatal("Gagal melakukan migrasi database: ", err)
	}

	fmt.Println("Berhasil terhubung ke database JapaneseClub dan migrasi tabel DataSiswa selesai!")
	
	DB = database

	BuatDataAdmin()
}

// buat data 
	func BuatDataAdmin() {
		var count int64
		DB.Model(&Admin{}).Count(&count)
		if count == 0 {
			admin := Admin{
				Username: "Admin1",
				Password: "ghulam159",
				NoHP:     "081234567890",
			}

			DB.Create(&admin)
			fmt.Println("Berhasil Membuat Data")
		}
	}
