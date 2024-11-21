package main

import (
	"SiAkademik/database"
	"SiAkademik/routes"
	"log"
)

func main() {

	// Hubungkan ke database
	database.Connect()

	// Lakukan migrasi tabel
	database.MigrateTables()

	// Jalankan server (contoh dengan Gin framework)

	log.Println("Server running at http://localhost:8080")
	routes.SetupRouter()

}
