package router

import (
	"orman-cicd-api/handler"
	"orman-cicd-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	users := api.Group("/users")
	users.Get("/:id", handler.GetUser)
	users.Post("/", handler.CreateUser)
	users.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	users.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Transaction
	transactions := api.Group("/transactions")
	transactions.Get("/", handler.GetAllTransactions)
	transactions.Get("/:id", handler.GetTransaction)
	transactions.Post("/", middleware.Protected(), handler.CreateTransaction)
	transactions.Delete("/:id", middleware.Protected(), handler.DeleteTransaction)

	// Company
	companies := api.Group("/companies")
	companies.Get("/", handler.GetAllCompanies)
	companies.Get("/:id", handler.GetCompany)
	companies.Post("/", middleware.Protected(), handler.CreateCompany)
	companies.Delete("/:id", middleware.Protected(), handler.DeleteCompany)

	// Company
	plantingAreas := api.Group("/plantingAreas")
	plantingAreas.Get("/", handler.GetAllPlantingAreas)
	plantingAreas.Get("/:id", handler.GetPlantingArea)
	plantingAreas.Post("/", middleware.Protected(), handler.CreatePlantingArea)
	plantingAreas.Delete("/:id", middleware.Protected(), handler.DeletePlantingArea)

	// Balance
	balances := api.Group("/balances")
	balances.Get("/", handler.GetAllBalances)
	balances.Get("/:id", handler.GetBalance)
	balances.Post("/", middleware.Protected(), handler.CreateBalance)
	balances.Delete("/:id", middleware.Protected(), handler.DeleteBalance)

}
