package matchingservice

import (
	"context"
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
	GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error)
}
type PresenceClient interface {
	GetPresence(ctx context.Context, request dto.GetPresenceRequest) (dto.GetPresenceResponse, error)
}
type Service struct {
	repo           Repo
	config         Config
	presenceClient PresenceClient
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
func (s Service) MatchWaitedUsers(ctx context.Context, req dto.MatchWaitedRequest) (dto.MatchWaitedResponse, error) {
	const op = "matchingservice.MatchWaitedUsers"
	list, err := s.repo.GetWaitingListByCategory(ctx, entity.FootballCategory)
	if err != nil {
		return dto.MatchWaitedResponse{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	for _, category := range entity.CategoryList() {

		for i := 0; i < len(list)-1; i = i + 2 {
			matchedList := entity.MatchedUsers{
				Category: category,
				UserID:   []uint{list[i].UserID, list[i+1].UserID},
			}

			// publish a new event

			//remove from waiting list
		}
	}
	return dto.MatchWaitedResponse{}, nil
}
