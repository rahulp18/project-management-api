package auth

import (
	"context"
	"errors"
	"project-management/internal/user"
)

type Service struct {
	userRepo *user.Repository
}

func NewService(repo *user.Repository) *Service {
	return &Service{
		userRepo: repo,
	}
}

func (service *Service) register(ctx context.Context, inputData user.UserInput) (*RegisterResponse, error) {
	if inputData.Email == "" || inputData.Name == "" || inputData.Password == "" {
		return nil, errors.New("Name email and password must required")
	}
	// hash password
	hashedPassword, err := HashedPassword(inputData.Password)
	if err != nil {
		return nil, err
	}
	inputData.Password = hashedPassword
	user, err := service.userRepo.CreateUser(ctx, inputData)
	if err != nil {
		return nil, err
	}
	// generate token from user
	token, err := GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}
	return &RegisterResponse{
		SessionToken: token,
	}, nil
}
