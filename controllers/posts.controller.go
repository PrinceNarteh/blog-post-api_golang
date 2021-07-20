package controllers

import (
	"fmt"

	"github.com/PrinceNarteh/blog-post-api_golang/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllPosts(c *fiber.Ctx) error {
	posts, err := services.GetAllPostsService()
	if err != nil {
		fmt.Println("error getting all posts")
	}
	return c.JSON(posts)
}

func GetPostById(c *fiber.Ctx) error {
	post, err := services.GetPostByIdService()
	if err != nil {
		fmt.Println("error getting post")
	}
	return c.JSON(post)
}

func CreatePost(c *fiber.Ctx) error {
	post, err := services.CreatePostService()
	if err != nil {
		fmt.Println("error creating post")
	}
	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	post, err := services.UpdatePostService()
	if err != nil {
		fmt.Println("error updating post")
	}
	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	post, err := services.GetPostByIdService()
	if err != nil {
		fmt.Println("error deleting post")
	}
	return c.JSON(post)
}
