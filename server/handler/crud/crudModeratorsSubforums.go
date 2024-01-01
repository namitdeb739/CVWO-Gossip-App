package crud

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)


func CreateModeratorsSubforums(c *fiber.Ctx) error {
	return CreateManyToManyEntry(c, model.ModeratorsSubforums{})
}

func GetAllModeratorsSubforums(c *fiber.Ctx) error {
	return GetAllManyToManyEntries(c, model.ModeratorsSubforums{})
}

func GetSingleModeratorsSubforums(c *fiber.Ctx) error {
	return GetSingleManyToManyEntry(c, model.ModeratorsSubforums{}, "UserID", "SubforumID")
}

func UpdateModeratorsSubforums(c *fiber.Ctx) error {
	return UpdateManyToManyEntry(c, model.ModeratorsSubforums{}, "UserID", "SubforumID")
}

func DeleteModeratorsSubforums(c *fiber.Ctx) error {
	return DeleteManyToManyEntry(c, model.ModeratorsSubforums{}, "UserID", "SubforumID")
}