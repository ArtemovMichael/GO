package accounts

import (
	"HW_2_fiber_flag/accounts/dto"
	"HW_2_fiber_flag/accounts/models"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c *fiber.Ctx) error {
	var request dto.CreateAccountRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	if len(request.Name) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Name is empty"})
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Account already exists"})
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Account created"})
}

func (h *Handler) GetAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Account not found"})
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(response)
}

func (h *Handler) DeleteAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	h.guard.Lock()

	if _, ok := h.accounts[name]; !ok {
		h.guard.Unlock()

		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Account not found"})
	}

	delete(h.accounts, name)

	h.guard.Unlock()

	return c.JSON(fiber.Map{"message": "Account deleted"})

}

func (h *Handler) UpdateAccountName(c *fiber.Ctx) error {
	old_name := c.Params("name")
	new_name := new(dto.ChangeAccountRequest)

	if err := c.BodyParser(&new_name); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	h.guard.Lock()

	if _, ok := h.accounts[old_name]; !ok {
		h.guard.Unlock()

		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Account not found"})
	}

	h.accounts[new_name.NewName] = h.accounts[old_name]
	h.accounts[new_name.NewName].Name = new_name.NewName
	delete(h.accounts, old_name)

	h.guard.Unlock()

	return c.JSON(fiber.Map{"message": "Account name updated"})
}

func (h *Handler) UpdateAccountAmount(c *fiber.Ctx) error {
	name := c.Params("name")
	new_amount := new(dto.ChangeAccountAmountRequest)

	if err := c.BodyParser(&new_amount); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	h.guard.Lock()

	if _, ok := h.accounts[name]; !ok {
		h.guard.Unlock()

		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Account not found"})
	}

	h.accounts[name].Amount = new_amount.Amount

	h.guard.Unlock()

	return c.JSON(fiber.Map{"message": "Account amount updated"})
}
