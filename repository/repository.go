package repository

import (
	"context"

	"github.com/mohamadafzal06/purl/entity"
)

type Repository interface {
	Save(context.Context, string, string, int64) error
	Load(context.Context, string) (string, error)
	LoadInfo(context.Context, string) (entity.URL, error)
	Close(context.Context) error
}
