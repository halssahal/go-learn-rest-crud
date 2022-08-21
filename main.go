package main

import (
	"learn/crud/models"
	"learn/crud/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run("localhost:1999")

}
