package accounts

import (
	"GO/Bank_account_Project/accounts/dto"
	"GO/Bank_account_Project/accounts/models"
	"sync"
	"log"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c *fiber.Ctx) error {
	var request dto.CreateAccountRequest

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)

		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	if len(request.Name) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.Status(fiber.StatusForbidden).SendString("Account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.JSON(h.accounts[request.Name])
}

func (h *Handler) GetAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	h.guard.RLock()
	account, ok := h.accounts[name]
	h.guard.RUnlock()

	if !ok {
		return c.Status(fiber.StatusNotFound).SendString("Account not found")
	}

	return c.JSON(account)
}

func (h *Handler) DeleteAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	h.guard.Lock()
	delete(h.accounts, name)
	h.guard.Unlock()

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) UpdateAccountName(c *fiber.Ctx) error {
	name := c.Params("name")
	var request dto.UpdateAccountNameRequest

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)

		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[name]; !ok {
		return c.Status(fiber.StatusNotFound).SendString("Account not found")
	}

	h.accounts[name].Name = request.Name

	return c.JSON(h.accounts[name])
}

func (h *Handler) UpdateAccountAmount(c *fiber.Ctx) error {
	name := c.Params("name")
	var request dto.UpdateAccountAmountRequest

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)

		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[name]; !ok {
		return c.Status(fiber.StatusNotFound).SendString("Account not found")
	}

	h.accounts[name].Amount = request.Amount

	return c.JSON(h.accounts[name])
}
