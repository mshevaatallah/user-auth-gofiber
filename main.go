package main

import (
	"belajar-fiber/database"
	"belajar-fiber/database/migrations"
	"belajar-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//initial func
	database.DatabaseInit()
	migrations.RunMigrations()
	routes.RouteInit(app)

	app.Listen(":8080")
}
