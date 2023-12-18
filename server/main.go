package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type user struct {
	User_ID string `json:"User_ID"`
}

var db *sql.DB

func main() {
	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	}) 

	go func() {
		log.Fatal(app.Listen(":4000"))
	}()

	

	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/gossip?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/users", getUser)
	router.POST("/users", createUser)

	router.Run(":8080")

}

func getUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.User_ID)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.Prepare(("INSERT INTO Users (User_ID) VALUES ($1)"))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newUser.User_ID); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newUser)
}

