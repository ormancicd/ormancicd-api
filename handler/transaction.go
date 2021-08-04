package handler

import (
	"orman-cicd-api/database"
	"orman-cicd-api/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllTransactions query all transactions
func GetAllTransactions(c *fiber.Ctx) error {
	db := database.DB
	var transactions []model.Transaction
	db.Find(&transactions)
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": transactions})
}

// GetTransaction query transaction
func GetTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var transaction model.Transaction
	db.Find(&transaction, id)
	if transaction.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No transaction found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Transaction found", "data": transaction})
}

// CreateTransaction new product
func CreateTransaction(c *fiber.Ctx) error {
	db := database.DB
	transaction := new(model.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}
	db.Create(&transaction)
	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": transaction})
}

// DeleteTransaction delete transaction
func DeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var transaction model.Transaction
	db.First(&transaction, id)
	if transaction.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}
	db.Delete(&transaction)
	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
