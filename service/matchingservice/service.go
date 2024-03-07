package matchingservice

import (
	"github.com/abbasfisal/game-app/deliver/dto"
	"github.com/abbasfisal/game-app/entity"
	"github.com/abbasfisal/game-app/pkg/richerror"
	"time"
)

type Config struct {
	WaitingTimeout time.Duration
}
type Repo interface {
	AddToWaitingList(userID uint, category entity.Category) error
}
type Service struct {
	repo   Repo
	config Config
}

func New() Service {
	return Service{}
}
func (s Service) AddToWaitingList(req dto.AddToWaitingListRequest) (dto.AddToWaitingListResponse, error) {
	const op = "matchingservice.AddToWaitingList"

	err := s.repo.AddToWaitingList(req.UserID, req.Category)
	if err != nil {
		return dto.AddToWaitingListResponse{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	//
	return dto.AddToWaitingListResponse{Timeout: s.config.WaitingTimeout}, nil

}
func (s Service) MatchWaitedUsers(req dto.MatchWaitedRequest) (dto.MatchWaitedResponse, error) {
	return dto.MatchWaitedResponse{}, nil
}
