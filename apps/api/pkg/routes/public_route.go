package routes

import (
	"go-playground/apps/api/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func to serve group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	route.Get("/states", controllers.GetStates)
}
