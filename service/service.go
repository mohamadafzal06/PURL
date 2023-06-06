package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/mohamadafzal06/purl/param"
	"github.com/mohamadafzal06/purl/pkg/randomstring"
	"github.com/mohamadafzal06/purl/repository"
)

type Service struct {
	repo repository.Repository
	rg   randomstring.RandomGenerator
}

func New(repo repository.Repository, rg randomstring.RandomGenerator) Service {
	return Service{
		repo: repo,
		rg:   rg,
	}
}

func (s Service) Short(ctx context.Context, sReq param.ShortRequest) (param.ShortResponse, error) {
	var response param.ShortResponse

	// short the requested url
	url := sReq.URL
	// TODO: check for Format of parsing
	expires, err := time.Parse("2006-01-02 15:04:05.728046 +0300 EEST", sReq.Expiry)
	if err != nil {
		return param.ShortResponse{}, fmt.Errorf("cannot parse the expiration time: %w\n", err)
	}

	key, err := s.repo.Save(ctx, url, expires)
	if err != nil {
		return param.ShortResponse{}, fmt.Errorf("cannot short the url: %w\n", err)
	}
	response.Key = key

	// return reposponse
	return response, nil
}

func (s Service) GetLong(ctx context.Context, lreq param.LongRequest) (param.LongResponse, error) {
	var response param.LongResponse
	// get the long-format from db
	url, err := s.repo.Load(ctx, lreq.Key)
	if err != nil {
		return param.LongResponse{}, fmt.Errorf("canntot retrieve the url by this key: %w\n", err)
	}

	// make the response
	response.LongURL = url

	return response, nil
}

func (s Service) GetLongInfo(ctx context.Context, lreq param.LongInfoRequest) (param.LongInfoResponse, error) {
	var response param.LongInfoResponse
	// LoadInfo(ctx context.Context, key string) (*entity.URL, error)

	shortLink, err := s.repo.LoadInfo(ctx, lreq.Key)
	if err != nil {
		return param.LongInfoResponse{}, fmt.Errorf("cannot retrieve url info with this key: %w\n", err)
	}
	response = param.LongInfoResponse{
		LongURL: shortLink.OriginalURL,
		Expiry:  shortLink.Expires,
		Visits:  strconv.Itoa(shortLink.Visits),
	}

	return response, nil
}
