package routes

import (
	"go-api/internal/auth"
	"go-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {

	app.Post("/qr", auth.JWTMiddleware(), handlers.QRHandler)

	app.Post("/qr/decompose", auth.JWTMiddleware(), handlers.QRDecomposeHandler)

}
