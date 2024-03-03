package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/abbasfisal/game-app/deliver/dto"
	"github.com/abbasfisal/game-app/entity"
	"github.com/abbasfisal/game-app/pkg/phonenumber"
	"github.com/abbasfisal/game-app/pkg/richerror"
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

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	//validate phone
	if !phonenumber.IsValid(req.PhoneNumber) {
		return dto.RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	//check uniqueness phone
	isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	if !isUnique {
		return dto.RegisterResponse{}, fmt.Errorf("phone number is not unique")
	}
	//validate name
	if len(req.Name) <= 3 {
		return dto.RegisterResponse{}, fmt.Errorf("name must be greater than 3")
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
		return dto.RegisterResponse{}, fmt.Errorf("unexpected eror %w", err)
	}

	//return use
	return dto.RegisterResponse{User: createdUser}, nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (s Service) GetProfile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	const op = "userservice.GetProfile"

	print("id", req.UserID)
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return dto.ProfileResponse{}, richerror.New(op).WithKind(richerror.KindNotFound).WithError(err).WithMeta(map[string]interface{}{"req": req})
	}

	return dto.ProfileResponse{Name: user.Name}, nil
}
