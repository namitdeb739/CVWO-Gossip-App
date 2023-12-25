package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/handler/crud"
	"github.com/namitdeb739/cvwo-gossip-app/handler/auth"
)

// Defines routes for the webapp
func SetupRoutes(app *fiber.App) {
	// Healthcheck for testing purposes
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")
	api.Post("/register", auth.Register)
	api.Post("/login", auth.Login)
	api.Get("/authuser", auth.AuthUser)
	api.Post("/logout", auth.Logout)

	user := api.Group("/user")
	user.Get("/", crud.GetAllUsers)
	user.Get("/:id", crud.GetSingleUser)
	user.Post("/", crud.CreateUser)
	user.Put("/:id", crud.UpdateUser)
	user.Delete("/:id", crud.DeleteUser)

	subforum := api.Group("/subforum")
	subforum.Get("/", crud.GetAllSubforums)
	subforum.Get("/:id", crud.GetSingleSubforum)
	subforum.Post("/", crud.CreateSubforum)
	subforum.Put("/:id", crud.UpdateSubforum)
	subforum.Delete("/:id", crud.DeleteSubforum)

	post := api.Group("/post")
	post.Get("/", crud.GetAllPosts)
	post.Get("/:id", crud.GetSinglePost)
	post.Post("/", crud.CreatePost)
	post.Put("/:id", crud.UpdatePost)
	post.Delete("/:id", crud.DeletePost)

	comment := api.Group("/comment")
	comment.Get("/", crud.GetAllComments)
	comment.Get("/:id", crud.GetSingleComment)
	comment.Post("/", crud.CreateComment)
	comment.Put("/:id", crud.UpdateComment)
	comment.Delete("/:id", crud.DeleteComment)

	tag := api.Group("/tag")
	tag.Get("/", crud.GetAllTags)
	tag.Get("/:id", crud.GetSingleTag)
	tag.Post("/", crud.CreateTag)
	tag.Put("/:id", crud.UpdateTag)
	tag.Delete("/:id", crud.DeleteTag)

	vote := api.Group("/vote")
	vote.Get("/", crud.GetAllVotes)
	vote.Get("/:id", crud.GetSingleVote)
	vote.Post("/", crud.CreateVote)
	vote.Put("/:id", crud.UpdateVote)
	vote.Delete("/:id", crud.DeleteVote)
}  