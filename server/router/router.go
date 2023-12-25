package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/handler/crud"
)

// Defines routes for the webapp
func SetupRoutes(app *fiber.App) {
	// Healthcheck for testing purposes
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")

	user := api.Group("/user")
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	subforum := api.Group("/subforum")
	subforum.Get("/", handler.GetAllSubforums)
	subforum.Get("/:id", handler.GetSingleSubforum)
	subforum.Post("/", handler.CreateSubforum)
	subforum.Put("/:id", handler.UpdateSubforum)
	subforum.Delete("/:id", handler.DeleteSubforum)

	post := api.Group("/post")
	post.Get("/", handler.GetAllPosts)
	post.Get("/:id", handler.GetSinglePost)
	post.Post("/", handler.CreatePost)
	post.Put("/:id", handler.UpdatePost)
	post.Delete("/:id", handler.DeletePost)

	comment := api.Group("/comment")
	comment.Get("/", handler.GetAllComments)
	comment.Get("/:id", handler.GetSingleComment)
	comment.Post("/", handler.CreateComment)
	comment.Put("/:id", handler.UpdateComment)
	comment.Delete("/:id", handler.DeleteComment)

	tag := api.Group("/tag")
	tag.Get("/", handler.GetAllTags)
	tag.Get("/:id", handler.GetSingleTag)
	tag.Post("/", handler.CreateTag)
	tag.Put("/:id", handler.UpdateTag)
	tag.Delete("/:id", handler.DeleteTag)

	vote := api.Group("/vote")
	vote.Get("/", handler.GetAllVotes)
	vote.Get("/:id", handler.GetSingleVote)
	vote.Post("/", handler.CreateVote)
	vote.Put("/:id", handler.UpdateVote)
	vote.Delete("/:id", handler.DeleteVote)
}  