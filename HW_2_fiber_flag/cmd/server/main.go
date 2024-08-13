package main

import (
	"HW_2_fiber_flag/accounts"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	handler := accounts.New()
	app := fiber.New()

	app.Post("/account", handler.CreateAccount)
	app.Get("/account/:name", handler.GetAccount)
	app.Get("/account/delete/:name", handler.DeleteAccount)
	app.Post("/account/change_name/:name", handler.UpdateAccountName)
	app.Post("/account/change_amount/:name", handler.UpdateAccountAmount)

	log.Fatal(app.Listen(":3000"))
}
