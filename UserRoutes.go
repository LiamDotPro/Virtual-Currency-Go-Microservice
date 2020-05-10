package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"time"
)

func UserRoutes(configureGroup *fiber.Group) {
	users := configureGroup.Group("/users")

	//Get
	users.Get("/acess", accessible)
	users.Get("/restricted", restricted)

	//Post
	users.Post("/login", login)

	//Delete

	//Patch

}

func login(c *fiber.Ctx) {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		c.SendStatus(fiber.StatusUnauthorized)
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) {
	c.Send("Accessible")
}

func restricted(c *fiber.Ctx) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	c.Send("Welcome " + name)
}
