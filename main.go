package main

import (
	"SiAkademik/database"
	"SiAkademik/routes"
	"log"
)

func main() {

	database.Connect()
	database.MigrateTables()
	log.Println("Server running at http://localhost:8080")
	routes.SetupRouter()

}
