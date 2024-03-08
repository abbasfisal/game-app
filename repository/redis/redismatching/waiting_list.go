package redismatching

import (
	"context"
	"fmt"
	"github.com/abbasfisal/game-app/entity"
	"github.com/abbasfisal/game-app/pkg/richerror"
	"github.com/abbasfisal/game-app/pkg/timestamp"
	"github.com/redis/go-redis/v9"
	"strconv"
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

func (d DB) GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error) {
	const op = "redismatching.GetWaitingListByCategory"

	min := fmt.Sprintf("%d", timestamp.Add(-2*time.Hour))
	max := strconv.Itoa(int(timestamp.Now()))

	list, err := d.adapter.Client().ZRangeByScoreWithScores(ctx, "a:b", &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: 0,
		Count:  0,
	}).Result()

	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	var result = make([]entity.WaitingMember, 0)
	for _, l := range list {
		userID, _ := strconv.Atoi(l.Member.(string))
		result = append(result, entity.WaitingMember{
			UserID:    uint(userID),
			Timestamp: int64(l.Score),
			Category:  category,
		})
	}
	return result, nil
}
