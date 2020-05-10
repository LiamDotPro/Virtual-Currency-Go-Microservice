package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/jwt"
)

// Load all middleware through here to clean main
func Middleware(app *fiber.App) {
	JWT(app)
}

func JWT(app *fiber.App) {
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
		Filter:     publicRoute,
	}))
}

func publicRoute(ctx *fiber.Ctx) bool {
	return true
}
