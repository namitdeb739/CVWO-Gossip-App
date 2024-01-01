package crud

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)


func CreatePostsTags(c *fiber.Ctx) error {
	return CreateManyToManyEntry(c, model.PostsTags{})
}

func GetAllPostsTags(c *fiber.Ctx) error {
	return GetAllManyToManyEntries(c, model.PostsTags{})
}

func GetSinglePostsTags(c *fiber.Ctx) error {
	return GetSingleManyToManyEntry(c, model.PostsTags{}, "PostID", "TagID")
}

func UpdatePostsTags(c *fiber.Ctx) error {
	return UpdateManyToManyEntry(c, model.PostsTags{}, "PostID", "TagID")
}

func DeletePostsTags(c *fiber.Ctx) error {
	return DeleteManyToManyEntry(c, model.PostsTags{}, "PostID", "TagID")
}