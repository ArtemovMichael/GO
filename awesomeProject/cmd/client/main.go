package main

import (
	"awesomeProject/accounts/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	Amount  int
	NewName string
}

func (cmd *Command) Do() error {
	switch cmd.Cmd {
	case "create":
		if err := cmd.Create(); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := cmd.Get(); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}
		return nil

	case "delete":
		if err := cmd.Delete(); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}
		return nil
	case "change_amount":
		if err := cmd.ChangeAmount(); err != nil {
			return fmt.Errorf("change amount failed: %w", err)
		}
		return nil
	case "change_name":
		if err := cmd.ChangeName(); err != nil {
			return fmt.Errorf("change name failed: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func (cmd *Command) Create() error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (cmd *Command) Get() error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account/%s", cmd.Host, cmd.Port, cmd.Name),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)
	fmt.Println()
	return nil
}

func (cmd *Command) Delete() error {

	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account/delete/%s", cmd.Host, cmd.Port, cmd.Name),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (cmd *Command) ChangeAmount() error {
	request := dto.ChangeAccountAmountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/change_amount/%s", cmd.Host, cmd.Port, cmd.Name),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (cmd *Command) ChangeName() error {
	request := dto.ChangeAccountRequest{
		Name:    cmd.Name,
		NewName: cmd.NewName,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/change_name/%s", cmd.Host, cmd.Port, cmd.Name),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newnameVal := flag.String("newname", "", "new name of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *amountVal,
		NewName: *newnameVal,
	}

	if err := cmd.Do(); err != nil {
		panic(err)
	}
}
