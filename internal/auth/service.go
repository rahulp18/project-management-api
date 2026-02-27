package auth

import (
	"context"
	"errors"
	"project-management/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepo *user.Repository
}

func NewService(repo *user.Repository) *Service {
	return &Service{
		userRepo: repo,
	}
}

func (service *Service) register(ctx context.Context, inputData user.UserInput) (*AuthResponse, error) {
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
	return &AuthResponse{
		SessionToken: token,
	}, nil
}

func (service *Service) login(ctx context.Context, inputData LoginRequest) (*AuthResponse, error) {
	if inputData.Email == "" || inputData.Password == "" {
		return nil, errors.New("Invalid data")
	}
	user, err := service.userRepo.FindByEmail(ctx, inputData.Email)

	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(inputData.Password),
	)

	if err != nil {
		return nil, errors.New("Invalid Email or Password")
	}
	// generate token
	token, err := GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{
		SessionToken: token,
	}, nil
}
