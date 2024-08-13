package main

import (
	"HW_3_grpc_flag/accounts/models"
	"HW_3_grpc_flag/proto"
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

func New() *server {
	return &server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type server struct {
	proto.BankAccountServiceServer
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (s *server) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	name := req.GetName()
	amount := req.GetAmount()

	if name == "" {
		return nil, errors.New("name is required")
	}

	s.guard.Lock()

	if _, err := s.accounts[name]; err {
		s.guard.Unlock()
		return nil, errors.New("account already exists")
	}

	s.accounts[name] = &models.Account{
		Name:   name,
		Amount: amount,
	}

	s.guard.Unlock()

	return &proto.CreateResponse{
		Name:   name,
		Amount: amount,
	}, nil
}

func (s *server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	name := req.GetName()

	if name == "" {
		return nil, errors.New("name is required")
	}

	s.guard.RLock()
	account, err := s.accounts[name]
	s.guard.RUnlock()

	if !err {
		return nil, errors.New("account not found")
	}

	return &proto.GetResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	name := req.GetName()

	if name == "" {
		return nil, errors.New("name is required")
	}

	s.guard.Lock()
	if _, err := s.accounts[name]; !err {
		s.guard.Unlock()
		return nil, errors.New("account not found")
	}

	delete(s.accounts, name)
	s.guard.Unlock()

	return &proto.DeleteResponse{
		Name: name,
	}, nil
}

func (s *server) UpdateName(ctx context.Context, req *proto.UpdateNameRequest) (*proto.UpdateNameResponse, error) {
	name := req.GetName()
	newName := req.GetNewName()

	if newName == "" {
		return nil, errors.New("name is required")
	}

	s.guard.Lock()
	if _, err := s.accounts[name]; !err {
		s.guard.Unlock()
		return nil, errors.New("account not found")
	}

	s.accounts[newName] = s.accounts[name]
	s.accounts[newName].Name = newName
	delete(s.accounts, name)
	s.guard.Unlock()

	return &proto.UpdateNameResponse{
		Name: newName,
	}, nil
}

func (s *server) UpdateAmount(ctx context.Context, req *proto.UpdateAmountRequest) (*proto.UpdateAmountResponse, error) {
	name := req.GetName()
	amount := req.GetAmount()

	s.guard.Lock()
	if _, err := s.accounts[name]; !err {
		s.guard.Unlock()
		return nil, errors.New("account not found")
	}

	s.accounts[name].Amount = amount
	s.guard.Unlock()

	return &proto.UpdateAmountResponse{
		Name:   name,
		Amount: amount,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		panic(err)
	}

	server := New()
	s := grpc.NewServer()

	proto.RegisterBankAccountServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
