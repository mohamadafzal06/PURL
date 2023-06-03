package service

import (
	"context"
	"fmt"

	"github.com/mohamadafzal06/purl/entity"
	"github.com/mohamadafzal06/purl/param"
	"github.com/mohamadafzal06/purl/pkg/randomstring"
)

type Service struct {
	repo Repository
	rg   randomstring.RandomGenerator
}

func New(repo Repository, rg randomstring.RandomGenerator) Service {
	return Service{
		repo: repo,
		rg:   rg,
	}
}

func (s Service) Short(ctx context.Context, sReq param.ShortRequest) (param.ShortResponse, error) {
	var response param.ShortResponse
	// check that the request is in db or not
	reqUrl := sReq.LongURL
	isexists, retrievKey, err := s.repo.IsURLInDB(ctx, reqUrl)

	if isexists {

		if err != nil {
			//TODO: check the error form repository layer
			return param.ShortResponse{}, fmt.Errorf("error while checking the url is in db or not: %w\n", err)
		}

		if retrievKey != "" {
			response.Key = retrievKey
			return response, nil
		}
	}

	// short the requested url
	generatedKey := s.rg.GenerateRandom()
	// TODO: is generatedKey is in db or not (check duplicaion)

	// insert the url and short-format in db
	insertedUrl := entity.URL{
		LongURL: sReq.LongURL,
		Key:     generatedKey,
	}

	_, err = s.repo.SetShortURL(ctx, insertedUrl)
	if err != nil {
		// TODO: checking all possible error
		return param.ShortResponse{}, fmt.Errorf("set the request in db failed: %w", err)
	}

	// TODO: checking for better return value
	response.Key = insertedUrl.Key

	// return reposponse
	return response, nil
}

func (s Service) GetLong(ctx context.Context, lreq param.LongRequest) (param.LongResponse, error) {
	var response param.LongResponse
	// get the long-format from db
	reqKey := lreq.Key
	isKeyExists, retrievURL, err := s.repo.IsKeyInDB(ctx, reqKey)

	if isKeyExists {
		if err != nil {
			//TODO: check the error form repository layer
			return param.LongResponse{}, fmt.Errorf("error while checking the key is in db or not: %w\n", err)
		}

		if retrievURL != "" {
			response.LongURL = retrievURL
		}
	}

	// make the response
	return response, nil
}
