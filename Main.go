package main

import (
	"github.com/gofiber/fiber"
)

// Version is the official version of the server app.
const Version = "v0.0.1"

// Location of the static content directory
const localStaticDir = "./static"

// @title Axual Notifications
// @version 1.0
// @description This is a simple notifications microservice created by Liam Read
// @contact.name Liam Read
// @contact.email liam@axual.com
// @license.name no license
// @host localhost:3645
// @BasePath /
func main() {

	// Load environment variables.
	env()

	// Make a connection to our database
	database()

	app := fiber.New()

	// All middleware is loaded before routes are added.
	Middleware(app)

	// Add Endpoints references
	Routes(app)

	_ = app.Listen(5000)
}
