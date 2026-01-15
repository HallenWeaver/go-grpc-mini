package server

import (
	"context"

	"github.com/HallenWeaver/go-grpc-mini/internal/service"
	userv1 "github.com/HallenWeaver/go-grpc-mini/proto/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	userv1.UnimplementedUserServiceServer
	store service.Store
}

func New(store service.Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	user, err := s.store.Create(service.User{
		Name:        req.GetName(),
		Email:       req.GetEmail(),
		DisplayName: req.GetDisplayName(),
	})

	if err != nil {
		return nil, err
	}

	return &userv1.CreateUserResponse{
		User: toProto(user),
	}, nil
}

func (s *Server) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	user, err := s.store.GetByID(req.GetId())
	if err != nil {
		return nil, err
	}

	return &userv1.GetUserResponse{
		User: toProto(user),
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.UpdateUserResponse, error) {
	user, err := s.store.Update(req.Id, &req.DisplayName)

	if err != nil {
		return nil, err
	}

	return &userv1.UpdateUserResponse{
		User: toProto(user),
	}, nil
}

func toProto(u service.User) *userv1.User {
	return &userv1.User{
		Id:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		IsActive:    u.Active,
		CreatedAt:   timestamppb.New(u.CreatedAt),
		UpdatedAt:   timestamppb.New(u.UpdatedAt),
	}
}
