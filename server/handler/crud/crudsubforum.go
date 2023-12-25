package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)



func CreateSubforum(c *fiber.Ctx) error {
	return CreateEntry(c, model.Subforum{})
}

func GetAllSubforums(c* fiber.Ctx) error {
	return GetAllEntries(c, model.Subforum{})
}

func GetSingleSubforum(c* fiber.Ctx) error {
	return GetSingleEntry(c, model.Subforum{}, "ID")
}

func UpdateSubforum(c *fiber.Ctx) error {
	return UpdateEntry(c, model.Subforum{}, "ID")
}

func DeleteSubforum(c* fiber.Ctx) error {
	return DeleteEntry(c, model.Subforum{}, "ID")
}