package param

type ShortRequest struct {
	LongURL string `json:"long_url"`
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
