package redispresence

import (
	"context"
	"github.com/abbasfisal/game-app/pkg/richerror"
	"time"
)

func (db DB) Upsert(ctx context.Context, key string, timestamp int64, expTime time.Duration) error {
	const op = "redispresence.Upsert"
	_, err := db.adapter.Client().Set(ctx, key, timestamp, expTime).Result()
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
