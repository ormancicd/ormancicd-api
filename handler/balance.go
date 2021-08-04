package handler

import (
	"github.com/gofiber/fiber/v2"
	"orman-cicd-api/database"
	"orman-cicd-api/model"
)

// GetAllBalances query all balances
func GetAllBalances(c *fiber.Ctx) error {
	db := database.DB
	var balances []model.Balance
	db.Find(&balances)
	return c.JSON(fiber.Map{"status": "success", "message": "All balances", "data": balances})
}

// GetBalance query balance
func GetBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var balance model.Balance
	db.Find(&balance, id)
	if balance.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No balance found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Balance found", "data": balance})
}

// CreateBalance new balance
func CreateBalance(c *fiber.Ctx) error {
	db := database.DB
	balance := new(model.Balance)
	if err := c.BodyParser(balance); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create balance", "data": err})
	}
	db.Create(&balance)
	return c.JSON(fiber.Map{"status": "success", "message": "Created balance", "data": balance})
}

// DeleteBalance delete balance
func DeleteBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var balance model.Balance
	db.First(&balance, id)
	if balance.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No balance found with ID", "data": nil})
	}
	db.Delete(&balance)
	return c.JSON(fiber.Map{"status": "success", "message": "Balance successfully deleted", "data": nil})
}


