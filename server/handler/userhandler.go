package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)


func CreateUser(c *fiber.Ctx) error {
	return CreateEntry(c, model.User{})
}

func GetAllUsers(c* fiber.Ctx) error {
	return GetAllEntries(c, model.User{})
}

func GetSingleUser(c* fiber.Ctx) error {
	return GetSingleEntry(c, model.User{}, "Username")
}

func UpdateUser(c *fiber.Ctx) error {
	return UpdateEntry(c, model.User{}, "Username")
}

func DeleteUser(c* fiber.Ctx) error {
	return DeleteEntry(c, model.User{}, "Username")
}