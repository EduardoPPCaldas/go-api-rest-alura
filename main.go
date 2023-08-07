package main

import (
	"fmt"

	"github.com/EduardoPPCaldas/go-api-rest-alura/database"
	"github.com/EduardoPPCaldas/go-api-rest-alura/models"
	"github.com/EduardoPPCaldas/go-api-rest-alura/routes"
)

func main() {
	db := database.GetDatabase()

	db.AutoMigrate(&models.Personality{})

	db.Create(&models.Personality{Id: 1, Name: "Nome", History: "History"})
	db.Create(&models.Personality{Id: 2, Name: "Nome2", History: "History"})
	db.Create(&models.Personality{Id: 3, Name: "Nome3", History: "History"})

	fmt.Println("Initializing server rest with Go")
	routes.HandleRequest()
}
