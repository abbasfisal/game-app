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
	GetUserByID(userID uint) (entity.User, error)
}

//type AuthParser interface {
//	ParseToken(bearerToken string) (*Claims, error)
//}

type AuthGenerator interface {
	CreateAccessToken(u entity.User) (string, error)
	CreateRefreshToken(u entity.User) (string, error)
}

// Service
type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{
		auth: authGenerator,
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
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type UserInfo struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
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

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error while generating jwt : %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)

	return LoginResponse{
		User: UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Tokens: Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

type ProfileRequest struct {
	UserID uint
}
type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) GetProfile(req ProfileRequest) (ProfileResponse, error) {
	print("id", req.UserID)
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	return ProfileResponse{Name: user.Name}, nil
}
