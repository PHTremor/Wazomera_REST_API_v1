package routes

import (
	"wazomera/app/controllers"
	"wazomera/core/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method: /Create Controllers
	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook) // create a new book

	// Routes for PUT method: /Update Controllers
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update one book by ID

	// Routes for DELETE method: /Delete Controllers
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID

}
