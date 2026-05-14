package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"backend-1/config"
	"backend-1/database"
	"backend-1/routes"
)

func main(){
	fmt.Println("Memulai Server...")

	fmt.Println("Server Berjalan di http://localhost:8000")
	router.run(":8000")
}
