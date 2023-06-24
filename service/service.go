package service

import (
	"context"
	"fmt"
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
	t := time.Now().Add(time.Duration(sReq.Expiry) * time.Hour)

	// Format the time as a string
	timeStr := t.Format(time.RFC3339)
	// Parse the string back into a time.Time value
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return param.ShortResponse{}, fmt.Errorf("Error parsing time: %w\n", err)
	}

	// Convert the time.Time to Unix timestamp
	unixTimestamp := parsedTime.Unix()

	key := s.rg.RandomString()

	err = s.repo.Save(ctx, key, url, unixTimestamp)
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
		Visits:  shortLink.Visits,
	}

	return response, nil
}
