package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	a := App{}
	a.Initialise(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	
	a.Run(":8010")
}


/*
fmt.Print("CVWO Gossip App")

	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	go func() {
		log.Fatal(app.Listen(":4000"))
	}()
*/