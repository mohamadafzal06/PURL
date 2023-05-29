package repository

import (
	"context"

	"github.com/mohamadafzal06/purl/entity"
)

type Purl interface {
	GetLongURL(ctx context.Context, surl entity.URL) (entity.URL, error)
	IsURLExist(ctx context.Context, url string) (entity.URL, error)
	IsKeyExist(ctx context.Context, key string) (entity.URL, error)
	SetShortURL(ctx context.Context, lurl entity.URL) (entity.URL, error)
}
