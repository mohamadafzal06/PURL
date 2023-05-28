package repository

import "context"

type Purl interface {
	GetLongURL(ctx context.Context, surl string) (string, error)
	SetShortURL(ctx context.Context, lurl string) (string, error)
}
