package repository

import (
	"context"

	"github.com/mohamadafzal06/purl/entity"
)

type Purl interface {
	GetLongURL(ctx context.Context, surl entity.URL) (entity.URL, error)
	SetShortURL(ctx context.Context, lurl entity.URL) (entity.URL, error)
}
