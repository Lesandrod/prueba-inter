package main

import (
	"go-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	routes.Register(app)

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
