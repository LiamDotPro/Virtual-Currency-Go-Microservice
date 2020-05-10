package main

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber"
)

func Routes(app *fiber.App) {

	group := getRoutingConfiguration(app)

	UserRoutes(group)
}

// Routing configuration
func getRoutingConfiguration(app *fiber.App) *fiber.Group {
	// Create default grouping
	group, err := getConfiguredGroupByVersion("v1", app)

	if err != nil {
		panic("The routing could not be configured")
	}

	return group
}

func getConfiguredGroupByVersion(versionIdentifier string, app *fiber.App) (*fiber.Group, error) {

	api := app.Group("/api")

	switch versionIdentifier {

	case "v1":
		return api.Group(fmt.Sprintf("/%v", versionIdentifier)), nil
	}

	return nil, errors.New("no version was matched when configuring routing")
}
