package userservice

import (
	"fmt"
	"github.com/abbasfisal/game-app/entity"
	"github.com/abbasfisal/game-app/pkg/phonenumber"
)

// Repository
type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

// Service
type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

// RegisterRequest -- Request
type RegisterRequest struct {
	Name        string
	PhoneNumber string
}

// RegisterResponse -- Response
type RegisterResponse struct {
	User entity.User
}

//var (
//	PhoneNumberIsNotValid = errors.New("phone number is not valid")
//)

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	//validate phone
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	//check uniqueness phone
	isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	if !isUnique {
		return RegisterResponse{}, fmt.Errorf("phone number is not unique")
	}
	//validate name
	if len(req.Name) <= 3 {
		return RegisterResponse{}, fmt.Errorf("name must be greater than 3")
	}
	//store user into db
	u := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}
	createdUser, err := s.repo.Register(u)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected eror %w", err)
	}

	//return use
	return RegisterResponse{createdUser}, nil
}
