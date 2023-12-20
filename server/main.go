package main

import (
	/* "database/sql"
	"log"
	"net/http" */

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
 	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/namitdeb739/cvwo-gossip-app/database"
	"github.com/namitdeb739/cvwo-gossip-app/router"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	
	router.SetupRoutes(app)
	
	// Handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":8080")
}

/* 
var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/gossip?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/users", getUser)
	app.Post("/users", createUser)

	// router := gin.Default()
	// router.GET("/users", getUser)
	// router.POST("/users", createUser)

	log.Fatal(app.Listen(":8080"))
	// log.Fatal(router.Run(":8080"))

}

func getUser(c *fiber.Ctx) {
	c.Append("Content-Type", "application/json")

	rows, err := db.Query(`SELECT * FROM gossip."Users"`)
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

	c.JSON(http.StatusOK)
}

func createUser(c *fiber.Ctx) {
	var newUser user
	if err := c.BodyParser(&newUser); err != nil {
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

func getUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query(`SELECT * FROM gossip."Users"`)
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
} */