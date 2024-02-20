package userservice

import "github.com/abbasfisal/game-app/entity"

type Service struct {
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

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {

}
