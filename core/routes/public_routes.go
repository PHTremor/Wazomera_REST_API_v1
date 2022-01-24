package routes

import (
	"wazomera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func describes group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("api/v1")

	// Routes for GET methods/ Controllers:
	route.Get("/books", controllers.GetBooks)              // get list of all books
	route.Get("/book/:id", controllers.GetBook)            // get one book by ID
	route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens

}
