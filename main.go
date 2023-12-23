// main.go

package main

import (
	"projectapi/database"
	"projectapi/router"
)

func main() {
	// Inisialisasi database
	database.ConnectDB()
	defer database.CloseDB()

	// Inisialisasi router
	r := router.SetupRouter()

	// Menjalankan server pada port tertentu
	r.Run(":3306")
}
