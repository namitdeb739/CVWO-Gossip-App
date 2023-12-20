package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")

	user := api.Group("/user")
	user.Get("/", handler.GetAllUsers)
	user.Get("/:username", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Delete("/:username", handler.DeleteUserByUsername)
}  