package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)

func CreatePost(c *fiber.Ctx) error {
	return CreateEntry(c, model.Post{})
}

func GetAllPosts(c* fiber.Ctx) error {
	return GetAllEntries(c, model.Post{})
}

func GetSinglePost(c* fiber.Ctx) error {
	return GetSingleEntry(c, model.Post{}, "ID")
}

func UpdatePost(c *fiber.Ctx) error {
	return UpdateEntry(c, model.Post{}, "ID")
}

func DeletePost(c* fiber.Ctx) error {
	return DeleteEntry(c, model.Post{}, "ID")
}