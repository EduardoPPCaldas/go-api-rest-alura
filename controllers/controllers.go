package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/EduardoPPCaldas/go-api-rest-alura/database"
	"github.com/EduardoPPCaldas/go-api-rest-alura/models"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Home")
}

func AllPersonalities(c echo.Context) error {
	var personalities []models.Personality
	database.DB.Find(&personalities)

	return c.JSON(http.StatusOK, personalities)
}

func GetPersonalityByID(c echo.Context) error {
	stringId := c.Param("id")

	id, err := strconv.Atoi(stringId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var personality models.Personality
	result := database.DB.Find(&personality, id)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Could not find any personality with the Id %d", id))
	}

	return c.JSON(http.StatusOK, personality)
}

func CreatePersonality(c echo.Context) error {
	var personality models.Personality

	if err := c.Bind(&personality); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := database.DB.Create(&personality)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadGateway, result.Error.Error())
	}

	return c.JSON(http.StatusCreated, personality)
}

func DeletePersonality(c echo.Context) error {
	id := c.Param("id")
	var personality models.Personality
	database.DB.Delete(&personality, id)

	return c.NoContent(http.StatusNoContent)
}

func UpdatePersonality(c echo.Context) error {
	id := c.Param("id")
	var personality models.Personality

	database.DB.First(&personality, id)
	c.Bind(&personality)

	database.DB.Save(&personality)
	return c.JSON(http.StatusOK, personality)
}
