package crud

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)


func CreateUser(c *fiber.Ctx) error {
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
											"message": "Invalid input: " + err.Error(),
											"data": user})
	}
	
	err = model.IsValidUser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
											"message": "Invalid input: " + err.Error(),
											"data": user})
	}

	return CreateEntry(c, model.User{})
}

func GetAllUsers(c *fiber.Ctx) error {
	return GetAllEntries(c, model.User{})
}

func GetSomeUsers(c *fiber.Ctx, searchKeys map[string]string) error {
	return GetSomeEntries(c, model.User{}, searchKeys)
}

func GetSingleUser(c *fiber.Ctx) error {
	return GetSingleEntry(c, model.User{}, "ID")
}

func UpdateUser(c *fiber.Ctx) error {
	return UpdateEntry(c, model.User{}, "ID")
}

func DeleteUser(c *fiber.Ctx) error {
	return DeleteEntry(c, model.User{}, "ID")
}