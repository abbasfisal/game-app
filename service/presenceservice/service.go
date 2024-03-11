package presenceservice

import (
	"context"
	"fmt"
	"github.com/abbasfisal/game-app/delivery/dto"
	"github.com/abbasfisal/game-app/pkg/richerror"
	"time"
)

type Config struct {
	ExpirationTime time.Duration
	Prefix         string
}
type Repo interface {
	Upsert(ctx context.Context, key string, timestamp int64, expTime time.Duration) error
}
type Service struct {
	config Config
	repo   Repo
}

func New(config Config, repo Repo) Service {
	return Service{
		config: config,
		repo:   repo,
	}
}

func (s Service) Upsert(ctx context.Context, req dto.UpsertPresenceRequest) (dto.UpsertPresenceResponse, error) {
	const op = "presenceservice.Upsert"
	err := s.repo.Upsert(ctx, fmt.Sprintf("%s:%d", s.config.Prefix, req.UserID), req.Timestamp, s.config.ExpirationTime)
	if err != nil {
		return dto.UpsertPresenceResponse{}, richerror.New(op).WithError(err)
	}
	return dto.UpsertPresenceResponse{}, nil
}
