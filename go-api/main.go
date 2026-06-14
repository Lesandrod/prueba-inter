package main

import (
	"go-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	routes.Register(app)

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
