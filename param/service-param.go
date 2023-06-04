package param

import "time"

type ShortRequest struct {
	URL    string        `json:"url"`
	Expiry time.Duration `json:"expiry"`
}

type ShortResponse struct {
	URL             string        `json:"url"`
	Key             string        `json:"key"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

type LongRequest struct {
	Key string `json:"key"`
}
type LongResponse struct {
	LongURL string `json:"long_url"`
}
