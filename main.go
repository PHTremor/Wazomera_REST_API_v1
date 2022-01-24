package main

import (
	"github.com/gofiber/fiber/v2"
	 "wazomera/core/configuration"
    "wazomera/core/middleware"
    "wazomera/core/routes"
    "wazomera/core/utils"

    _ "github.com/joho/godotenv/autoload"
	_"wazomera/docs"
)

// @name Authorization
// @BasePath /api
func main()  {
	// Define Fiber config.
	config := configuration.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes
	 // Routes.
	 routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	 routes.PublicRoutes(app)  // Register a public routes for app.
	 routes.PrivateRoutes(app) // Register a private routes for app.
	 routes.NotFoundRoute(app)

	 // Start server (with graceful shutdown).
	 utils.StartServerWithGracefulShutdown(app)
}