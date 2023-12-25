package crud

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)

func CreateComment(c *fiber.Ctx) error {
	return CreateEntry(c, model.Comment{})
}

func GetAllComments(c* fiber.Ctx) error {
	return GetAllEntries(c, model.Comment{})
}

func GetSingleComment(c* fiber.Ctx) error {
	return GetSingleEntry(c, model.Comment{}, "ID")
}

func UpdateComment(c *fiber.Ctx) error {
	return UpdateEntry(c, model.Comment{}, "ID")
}

func DeleteComment(c* fiber.Ctx) error {
	return DeleteEntry(c, model.Comment{}, "ID")
}