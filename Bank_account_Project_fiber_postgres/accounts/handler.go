package accounts

import (
	"Bank_account_Project_fiber_postgres/accounts/dto"
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func CreateAccount(c *fiber.Ctx) error {
	var request dto.CreateAccountRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	if len(request.Name) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Name is empty"})
	}

	ctx := context.Background()

	db := dto.GetDB()
	if dto.IsAccountExists(request.Name, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account already exists"})
	}

	_, err := db.ExecContext(ctx, "INSERT INTO accounts (name, amount) VALUES ($1, $2)", request.Name, request.Amount)

	if err != nil {
		log.Fatal(err)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Account created"})
}

func GetAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	ctx := context.Background()

	db := dto.GetDB()
	if !dto.IsAccountExists(name, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account does not exist"})
	}

	rows, err := db.QueryContext(ctx, "SELECT name, amount FROM accounts WHERE name = $1", name)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	var account dto.GetAccountResponse

	for rows.Next() {
		if err := rows.Scan(&account.Name, &account.Amount); err != nil {
			log.Fatal(err)
		}
	}

	return c.JSON(account)
}

func DeleteAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	ctx := context.Background()

	db := dto.GetDB()
	if !dto.IsAccountExists(name, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account does not exist"})
	}

	_, err := db.ExecContext(ctx, "DELETE FROM accounts WHERE name = $1", name)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{"message": "Account deleted"})
}

func UpdateAccountName(c *fiber.Ctx) error {
	old_name := c.Params("name")
	new_name := new(dto.ChangeAccountRequest)

	if err := c.BodyParser(&new_name); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	ctx := context.Background()

	db := dto.GetDB()
	if !dto.IsAccountExists(old_name, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account does not exist"})
	}

	_, err := db.ExecContext(ctx, "UPDATE accounts SET name = $1 WHERE name = $2", new_name.NewName, old_name)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{"message": "Account name updated"})
}

func UpdateAccountAmount(c *fiber.Ctx) error {
	name := c.Params("name")
	new_amount := new(dto.ChangeAccountAmountRequest)

	if err := c.BodyParser(&new_amount); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	ctx := context.Background()

	db := dto.GetDB()
	if !dto.IsAccountExists(name, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account does not exist"})
	}

	_, err := db.ExecContext(ctx, "UPDATE accounts SET amount = $1 WHERE name = $2", new_amount.Amount, name)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{"message": "Account amount updated"})
}
