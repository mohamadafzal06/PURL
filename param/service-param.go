package param

type ShortRequest struct {
	LongURL string `json:"long_url"`
}
type ShortResponse struct {
	ShortURL string `json:"short_url"`
}

type LongRequest struct {
	ShortURL string `json:"short_url"`
}
type LongResponse struct {
	LongURL string `json:"long_url"`
}
