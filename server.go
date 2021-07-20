package main

import (
	"github.com/PrinceNarteh/blog-post-api_golang/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// initializing fiber app
	app := fiber.New()

	// setting up routes
	routers.PostsRoute(app)

	app.Listen(":4000")
}
