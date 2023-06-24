package param

import (
	"fmt"
	"time"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ShortRequest struct {
	URL string `json:"url"`
	// Number of Hours from Now
	Expiry int64 `json:"expiry"`
}

type ShortResponse struct {
	Key string `json:"key"`
}

type LongRequest struct {
	Key string `json:"key"`
}
type LongResponse struct {
	LongURL string `json:"long_url"`
}

type LongInfoRequest struct {
	Key string `json:"key"`
}
type LongInfoResponse struct {
	LongURL string `json:"long_url"`
	Expiry  int64  `json:"expiry"`
	Visits  int    `json:"visits"`
}

func (sr ShortRequest) ValidateShort() error {
	if err := validation.ValidateStruct(&sr,
		validation.Field(&sr.URL, validation.Required, is.RequestURI),
		validation.Field(&sr.Expiry, validation.Min(time.Now())),
	); err != nil {
		return fmt.Errorf("url request validation failed: %w", err)
	}

	return nil
}

func (lr LongRequest) ValidateLong() error {
	if err := validation.ValidateStruct(&lr,
		// TODO: add more validation for key field
		validation.Field(&lr.Key, validation.Required),
	); err != nil {
		return fmt.Errorf("key request validation failed: %w", err)
	}

	return nil
}

func (lr LongInfoRequest) ValidateLongInfo() error {
	if err := validation.ValidateStruct(&lr,
		// TODO: add more validation for key field
		validation.Field(&lr.Key, validation.Required),
	); err != nil {
		return fmt.Errorf("key request validation failed: %w", err)
	}

	return nil
}
