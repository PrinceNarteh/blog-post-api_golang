package routers

import (
	"github.com/PrinceNarteh/blog-post-api_golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func PostsRoute(app *fiber.App) {
	api := app.Group("/api/posts")
	api.Get("/", controllers.GetAllPosts)
	api.Post("/", controllers.CreatePost)
	api.Get("/:id", controllers.GetPostById)
	api.Patch("/:id", controllers.UpdatePost)
	api.Delete("/:id", controllers.DeletePost)
}
