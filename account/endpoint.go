package account

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser endpoint.Endpoint
	Login endpoint.Endpoint
}
func MakeEndpoints(s Service) Endpoints{
	return Endpoints{
		CreateUser: nil,
		GetUser:    nil,
		Login:      nil,
	}
}

func makeCreateUserEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok:ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		id, err := s.GetUser(ctx, req.Email)
		return GetUserResponse{Id:id}, err
	}
}

func makeLoginEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoginRequest)
		token, err := s.Login(ctx, req.Email, req.Password)
		return LoginResponse{Token:token}, err
	}
}