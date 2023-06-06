package entity

type URL struct {
	ID          uint64 `json:"id" redis:"id"`
	OriginalURL string `json:"original_url" redis:"original_url"`
	Expires     string `json:"expires" redis:"expires"`
	Visits      int    `json:"visits" redis:"visits"`
}
