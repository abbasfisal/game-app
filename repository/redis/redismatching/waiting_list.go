package redismatching

import (
	"context"
	"github.com/abbasfisal/game-app/entity"
	"github.com/abbasfisal/game-app/pkg/richerror"
	"github.com/redis/go-redis/v9"
	"time"
)

func (d DB) AddToWaitingList(userID uint, category entity.Category) error {
	const op = "redismatching.AddTOWaitingList"

	_, err := d.adapter.Client().ZAdd(context.Background(), "a:b", redis.Z{
		Score:  float64(time.Now().UnixMicro()),
		Member: userID,
	}).Result()
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
