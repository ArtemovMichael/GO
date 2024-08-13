package main

import (

	"fmt"
	"log"

	"HW_4_fiber_postgres/accounts"
	"HW_4_fiber_postgres/accounts/dto"
	
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := dto.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/account", accounts.CreateAccount)
	app.Get("/account/:name", accounts.GetAccount)
	app.Get("/account/delete/:name", accounts.DeleteAccount)
	app.Post("/account/change_name/:name", accounts.UpdateAccountName)
	app.Post("/account/change_amount/:name", accounts.UpdateAccountAmount)

	fmt.Println("Server is running on port 3000")

	log.Fatal(app.Listen(":3000"))

}
