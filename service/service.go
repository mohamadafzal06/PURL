package service

import (
	"context"
	"fmt"

	"github.com/mohamadafzal06/purl/entity"
	"github.com/mohamadafzal06/purl/param"
	"github.com/mohamadafzal06/purl/repository"
)

type Service struct {
	repo        repository.Purl
	shortenFunc func(int) string
}

func New(repo repository.Purl, sFunc func(int) string) Service {
	return Service{
		repo:        repo,
		shortenFunc: sFunc,
	}
}

func (s *Service) Short(ctx context.Context, sReq param.ShortRequest) (param.ShortResponse, error) {
	var response param.ShortResponse
	// check that the request is in db or not
	url := sReq.LongURL
	retrievURL, err := s.repo.IsURLExist(ctx, url)
	if err != nil {
		//TODO: check the error form repository layer
		return param.ShortResponse{}, fmt.Errorf("error while checking the url is in db or not: %w\n", err)
	}
	if retrievURL.Key != "" {
		response.ShortURL = retrievURL.Key
		return response, nil
	}

	// short the requested url
	generatedKey := s.shortenFunc(6)

	// insert the url and short-format in db
	insertedUrl := entity.URL{
		LURL: sReq.LongURL,
		Key:  generatedKey,
	}
	resultURL, err := s.repo.SetShortURL(ctx, insertedUrl)
	if err != nil {
		// TODO: checking all possible error
		return param.ShortResponse{}, fmt.Errorf("set the request in db failed: %w", err)
	}

	// TODO: checking for better return value
	response.ShortURL = resultURL.Key

	// return reposponse
	return response, nil
}

func (s *Service) GetLong(ctx context.Context, surl param.LongRequest) (param.LongResponse, error) {
	var response param.LongResponse
	// get the long-format from db
	reqKey := surl.ShortURL
	retrievURL, err := s.repo.IsKeyExist(ctx, reqKey)
	if err != nil {
		//TODO: check the error form repository layer
		return param.LongResponse{}, fmt.Errorf("error while checking the key is in db or not: %w\n", err)
	}
	if retrievURL.LURL != "" {
		response.LongURL = retrievURL.LURL
	}

	// make the response
	return response, nil
}
