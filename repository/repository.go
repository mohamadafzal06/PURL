package repository

import (
	"context"
	"time"

	"github.com/mohamadafzal06/purl/entity"
)

type Repository interface {
	Save(context.Context, string, time.Time) (string, error)
	Load(context.Context, string) (string, error)
	LoadInfo(context.Context, string) (*entity.URL, error)
	Close(context.Context) error
}
