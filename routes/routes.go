package routes

import (
	"log"

	"github.com/EduardoPPCaldas/go-api-rest-alura/controllers"
	"github.com/labstack/echo/v4"
)

func HandleRequest() {
	e := echo.New()

	e.GET("/", controllers.Home)
	e.GET("/api/personalities", controllers.AllPersonalities)
	e.GET("/api/personalities/:id", controllers.GetPersonalityByID)
	e.POST("/api/personalities", controllers.CreatePersonality)
	e.DELETE("/api/personalities/:id", controllers.DeletePersonality)
	e.PUT("/api/personalities/:id", controllers.UpdatePersonality)

	log.Fatal(e.Start(":1323"))
}
