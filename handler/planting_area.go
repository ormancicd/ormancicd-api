package handler

import (
	"github.com/gofiber/fiber/v2"
	"orman-cicd-api/database"
	"orman-cicd-api/model"
)

// GetAllPlantingAreas query all plantingAreas
func GetAllPlantingAreas(c *fiber.Ctx) error {
	db := database.DB
	var plantingAreas []model.PlantingArea
	db.Find(&plantingAreas)
	return c.JSON(fiber.Map{"status": "success", "message": "All plantingAreas", "data": plantingAreas})
}

// GetPlantingArea query plantingArea
func GetPlantingArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var plantingArea model.PlantingArea
	db.Find(&plantingArea, id)
	if plantingArea.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No plantingArea found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "PlantingArea found", "data": plantingArea})
}

// CreatePlantingArea new plantingArea
func CreatePlantingArea(c *fiber.Ctx) error {
	db := database.DB
	plantingArea := new(model.PlantingArea)
	if err := c.BodyParser(plantingArea); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create plantingArea", "data": err})
	}
	db.Create(&plantingArea)
	return c.JSON(fiber.Map{"status": "success", "message": "Created plantingArea", "data": plantingArea})
}

// DeletePlantingArea delete plantingArea
func DeletePlantingArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var plantingArea model.PlantingArea
	db.First(&plantingArea, id)
	if plantingArea.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No plantingArea found with ID", "data": nil})
	}
	db.Delete(&plantingArea)
	return c.JSON(fiber.Map{"status": "success", "message": "PlantingArea successfully deleted", "data": nil})
}

