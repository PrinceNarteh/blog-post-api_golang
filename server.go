package main

import (
	"github.com/PrinceNarteh/blog-post-api_golang/config"
	"github.com/PrinceNarteh/blog-post-api_golang/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// connecting to database
	config.Setup()

	// initializing fiber app
	app := fiber.New()

	// setting up routes
	routers.PostsRoute(app)

	app.Listen(":4000")
}
