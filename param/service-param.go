package param

type ShortRequest struct {
	URL    string `json:"url"`
	Expiry string `json:"expiry"`
}

type ShortResponse struct {
	URL             string `json:"url"`
	Key             string `json:"key"`
	Expiry          string `json:"expiry"`
	XRateRemaining  int    `json:"rate_limit"`
	XRateLimitReset string `json:"rate_limit_reset"`
}

type LongRequest struct {
	Key string `json:"key"`
}
type LongResponse struct {
	LongURL string `json:"long_url"`
}
