package userservice

import (
	"fmt"
	"github.com/abbasfisal/game-app/delivery/dto"
	"github.com/abbasfisal/game-app/pkg/richerror"
)

func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	const op = "userservice.Login"
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).WithError(err)

	}

	if !exist {
		return dto.LoginResponse{}, fmt.Errorf("username or password isnt correct! ")
	}
	if user.Password != GetMD5Hash(req.Password) {
		return dto.LoginResponse{}, fmt.Errorf("username or password isnt correct! ")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error while generating jwt : %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)

	return dto.LoginResponse{
		User: dto.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Tokens: dto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
