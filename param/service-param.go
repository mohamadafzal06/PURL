package param

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
