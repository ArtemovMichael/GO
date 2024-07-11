package main

import (
	"flag"
	"fmt"
	"time"

	"context"

	"Bank_account_Project_grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Command struct {
	Port    string
	Host    string
	Cmd     string
	Name    string
	Amount  float64
	NewName string
}

func (cmd *Command) Do() error {
	conn, err := grpc.NewClient(cmd.Host+":"+cmd.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	switch cmd.Cmd {
	case "create":
		if err := cmd.Create(conn); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}
		return nil
	case "get":
		if err := cmd.Get(conn); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}
		return nil

	case "delete":
		if err := cmd.Delete(conn); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}
		return nil
	case "change_amount":
		if err := cmd.ChangeAmount(conn); err != nil {
			return fmt.Errorf("change amount failed: %w", err)
		}
		return nil
	case "change_name":
		if err := cmd.ChangeName(conn); err != nil {
			return fmt.Errorf("change name failed: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func (cmd *Command) Create(conn *grpc.ClientConn) error {

	c := proto.NewBankAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Create(ctx, &proto.CreateRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Account created: %s\n", res.GetName())

	return nil
}

func (cmd *Command) Get(conn *grpc.ClientConn) error {

	c := proto.NewBankAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Get(ctx, &proto.GetRequest{
		Name: cmd.Name,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Account: %s, Amount: %f\n", res.GetName(), res.GetAmount())

	return nil
}

func (cmd *Command) Delete(conn *grpc.ClientConn) error {

	c := proto.NewBankAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Delete(ctx, &proto.DeleteRequest{
		Name: cmd.Name,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Account deleted: %s\n", res.GetName())

	return nil
}

func (cmd *Command) ChangeAmount(conn *grpc.ClientConn) error {

	c := proto.NewBankAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.UpdateAmount(ctx, &proto.UpdateAmountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Account: %s, Amount: %f\n", res.GetName(), res.GetAmount())

	return nil
}

func (cmd *Command) ChangeName(conn *grpc.ClientConn) error {

	c := proto.NewBankAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.UpdateName(ctx, &proto.UpdateNameRequest{
		Name:    cmd.Name,
		NewName: cmd.NewName,
	})

	if err != nil {
		return err
	}

	fmt.Printf("New name: %s\n", res.GetName())

	return nil
}

func main() {
	portVal := flag.String("port", "3000", "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Float64("amount", 0, "amount of account")
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
