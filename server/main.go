package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)



func main() {
	fmt.Print("CVWO Gossip App")

	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":4000"))

	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/mydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	
}