package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/abbasfisal/game-app/entity"
	"github.com/abbasfisal/game-app/pkg/phonenumber"
)

// Repository
type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
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
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
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
		Password:    GetMD5Hash(req.Password),
	}
	createdUser, err := s.repo.Register(u)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected eror %w", err)
	}

	//return use
	return RegisterResponse{createdUser}, nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type LoginResponse struct {
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {

	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %v", err)
	}

	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password isnt correct! ")
	}
	if user.Password != GetMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password isnt correct! ")
	}

	return LoginResponse{}, nil
}
