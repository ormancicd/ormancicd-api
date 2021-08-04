package handler

import (
	"github.com/gofiber/fiber/v2"
	"orman-cicd-api/database"
	"orman-cicd-api/model"
)

// GetAllCompanies query all companies
func GetAllCompanies(c *fiber.Ctx) error {
	db := database.DB
	var companies []model.Company
	db.Find(&companies)
	return c.JSON(fiber.Map{"status": "success", "message": "All companies", "data": companies})
}

// GetCompany query company
func GetCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var company model.Company
	db.Find(&company, id)
	if company.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No company found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Company found", "data": company})
}

// CreateCompany new company
func CreateCompany(c *fiber.Ctx) error {
	db := database.DB
	company := new(model.Company)
	if err := c.BodyParser(company); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create company", "data": err})
	}
	db.Create(&company)
	return c.JSON(fiber.Map{"status": "success", "message": "Created company", "data": company})
}

// DeleteCompany delete company
func DeleteCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var company model.Company
	db.First(&company, id)
	if company.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No company found with ID", "data": nil})
	}
	db.Delete(&company)
	return c.JSON(fiber.Map{"status": "success", "message": "Company successfully deleted", "data": nil})
}

