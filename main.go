package main

import (
	"log"
	"orman-cicd-api/database"
	"orman-cicd-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

	defer database.DB.Close()
}
