package service

import (
	"context"

	"github.com/mohamadafzal06/purl/entity"
)

type Repository interface {
	GetLongURL(ctx context.Context, key string) (string, error)
	IsURLInDB(ctx context.Context, url string) (bool, string, error)
	IsKeyInDB(ctx context.Context, key string) (bool, string, error)
	SetShortURL(ctx context.Context, lurl entity.URL) (uint64, error)
}
