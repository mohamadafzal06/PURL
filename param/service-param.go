package param

type ShortRequest struct {
	URL    string `json:"url"`
	Expiry string `json:"expiry"`
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
	Expiry  string `json:"expiry"`
	Visits  string `json:"visits"`
}
