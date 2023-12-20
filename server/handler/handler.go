package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/database"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
											"message": "Invalid input",
											"data": "err"})
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
											"message": "Could not create user",
											"data": "err"})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success",
										"message": "User created",
										"data": user})
}

func GetAllUsers(c* fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Users not found",
											"data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": "Users found",
										"data": users})
}

func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db

	username := c.Params("Username")
	var user model.User

	search := db.Where("Username = ?", username).First(&user)
	if search.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "User " + username +  " not found",
											"data": nil})
	}
	

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": "User found",
										"data": username})
}

/* func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		User_ID string `JSON:"User_ID"`
	}

	db := database.DB.Db

	var user model.User

	user_id := c.Params("User_ID")

	search := db.Where("User_ID = ?", user_id).First(&user)

	if search.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "User " + user_id +  " not found",
											"data": nil})
	}

	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Invalid input",
											"data": err})
	}

	user.User_ID = updateUserData.User_ID

	db.Save(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": "User found",
										"data": user})
} */

func DeleteUserByUsername(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User

	username := c.Params("Username")

	search := db.Where("Username = ?", username).First(&user)
	if search.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "User " + username +  " not found",
											"data": nil})
	}

	err := db.Delete(&user, "Username = ?", username).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Failed to delete user",
											"data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "error",
										"message": "User " + username +  " deleted"})
}