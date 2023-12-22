package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)

func CreateTag(c *fiber.Ctx) error {
	return CreateEntry(c, model.Tag{})
}

func GetAllTags(c* fiber.Ctx) error {
	return GetAllEntries(c, model.Tag{})
}

func GetSingleTag(c* fiber.Ctx) error {
	return GetSingleEntry(c, model.Tag{}, "ID")
}

func UpdateTag(c *fiber.Ctx) error {
	return UpdateEntry(c, model.Tag{}, "ID")
}

func DeleteTag(c* fiber.Ctx) error {
	return DeleteEntry(c, model.Tag{}, "ID")
}