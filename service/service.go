package service

import (
	"github.com/mohamadafzal06/purl/param"
	"github.com/mohamadafzal06/purl/repository"
)

type Service struct {
	repo        repository.Purl
	shortenFunc func(int) string
}

func New(repo repository.Purl, sFunc func(int) string) *Service {
	return &Service{
		repo:        repo,
		shortenFunc: sFunc,
	}
}

func (s *Service) Short(sReq param.ShortRequest) (param.ShortResponse, error) {
	// check that the request is in db or not

	// short the requested url
	surl := s.shortenFunc(6)
	resp := param.ShortResponse{
		ShortURL: surl,
	}

	// insert the url and short-format in db

	// return reposponse
	return resp, nil
}

func (s *Service) GetLong(surl param.LongRequest) (param.LongResponse, error) {
	// get the long-format from db

	// make the response

	return param.LongResponse{}, nil
}
