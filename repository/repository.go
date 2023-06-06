package repository

import (
	"time"

	"github.com/mohamadafzal06/purl/entity"
)

type Repository interface {
	Save(string, time.Time) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*entity.URL, error)
	Close() error
}
