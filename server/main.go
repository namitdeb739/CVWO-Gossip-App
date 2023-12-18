package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Print("CVWO Gossip App")

	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	var err error
	connectionStr := "postgres://postgres:postgres@localhost:5432/gossip?sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", createUsers)

	router.Run("localhost:8080")

	go func() {
		log.Fatal(app.Listen(":4000"))
	}()
	
}

func getUsers (c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.scan(&u.User_ID)
		if err != nil {
			log.Fatal(err)
		}
		users.append(users, u)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, users)
}