package service

import "github.com/mohamadafzal06/purl/repository"

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

type ShortRequest struct {
	LongURL string
}
type ShortResponse struct {
	ShortURL string
}

type LongRequest struct {
	ShortURL string
}
type LongResponse struct {
	LongURL string
}

func (s *Service) Short(sReq ShortRequest) (ShortResponse, error) {
	// check that the request is in db or not

	// short the requested url
	surl := s.shortenFunc(6)
	resp := ShortResponse{
		ShortURL: surl,
	}

	// insert the url and short-format in db

	// return reposponse
	return resp, nil
}
func (s *Service) GetLong(surl LongRequest) (LongResponse, error) {
	// get the long-format from db

	// make the response

	return LongResponse{}, nil
}
