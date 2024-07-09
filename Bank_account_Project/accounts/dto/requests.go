package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type GetAccountRequest struct {
	Name string `json:"name"`
}

type UpdateAccountNameRequest struct {
	Name string `json:"name"`
}

type UpdateAccountAmountRequest struct {
	Amount int `json:"amount"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}