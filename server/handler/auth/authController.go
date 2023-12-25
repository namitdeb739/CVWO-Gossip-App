package auth

import (
	"reflect"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/database"
	"github.com/namitdeb739/cvwo-gossip-app/handler/crud"
	"github.com/namitdeb739/cvwo-gossip-app/model"
)

const Secretkey = "secret"

func Register(c *fiber.Ctx) error {
	var entry map[string]string

	if err := c.BodyParser(&entry); err != nil {
		return err
	}

	if len(entry["password"]) < 8 {
		return c.Status(422).JSON(fiber.Map{"status": "error",
											"message": "Invalid Password: Length must be at least 8 characters",
											"data": nil})
	}
	
	return crud.CreateUser(c)
}

func Login(c *fiber.Ctx) error {
	db := database.DB.Db

	var entry map[string]string
	if err := c.BodyParser(&entry)
	err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error",
											"message": "Invalid input: " + err.Error(),
											"data": nil})
	}

	var user model.User

	search := db.Where("Username = ?", entry["username"]).First(&user)
	if search.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "User" + user.Username + " not found",
											"data": nil})
	}

	if user.Password != entry["password"] {
		return c.Status(401).JSON(fiber.Map{"status": "error",
											"message": "Incorrect password",
											"data": nil})
	}
	
	// Preload all associations of the entry (show relations)
	val := reflect.ValueOf(&user).Elem()
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Struct {
            associationName := val.Type().Field(i).Name
            db.Preload(associationName).Find(entry)
        }
    }

	claims :=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Expires in 1 day
	})

	token, err := claims.SignedString([]byte(Secretkey))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"status": "error",
											"message": "Could not log in",
											"data": nil})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": "Logged in",
										"data": token})
}

func AuthUser(c *fiber.Ctx) error {
	db := database.DB.Db

	var user model.User
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"status": "error",
											"message": "Unauthenticated",
											"data": nil})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	db.Where("ID = ?", claims.Issuer).First(&user)
	
	// Preload all associations of the entry (show relations)
	val := reflect.ValueOf(&user).Elem()
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Struct {
            associationName := val.Type().Field(i).Name
            db.Preload(associationName).Find(&user)
        }
    }

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": "User authenticated",
										"data": user})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": "User logged out"})
}