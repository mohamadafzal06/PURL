package param

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
