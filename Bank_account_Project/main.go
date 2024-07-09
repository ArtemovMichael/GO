package main

import (
	"log"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Post("/account", CreateAccount)
	app.Get("/account/:name", GetAccount)
	app.Delete("/account/:name", DeleteAccount)
	app.Put("/account/:name", UpdateAccountName)
	app.Put("/account/:name/amount", UpdateAccountAmount)

	// Start the server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
