package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	v1.Get("/", handler.GetAllUsers)
	v1.Get("/:user_id", handler.GetSingleUser)
	v1.Post("/", handler.CreateUser)
	// v1.Put("/:user_id", handler.UpdateUser)
	v1.Delete("/:user_id", handler.DeleteUserByUserID)
}