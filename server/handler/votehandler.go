package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)



func CreateVote(c *fiber.Ctx) error {
	return CreateEntry(c, model.Vote{})
}

func GetAllVotes(c* fiber.Ctx) error {
	return GetAllEntries(c, model.Vote{})
}

func GetSingleVote(c* fiber.Ctx) error {
	return GetSingleEntry(c, model.Vote{}, "ID")
}

func UpdateVote(c *fiber.Ctx) error {
	return UpdateEntry(c, model.Vote{}, "ID")
}

func DeleteVote(c* fiber.Ctx) error {
	return DeleteEntry(c, model.Vote{}, "ID")
}