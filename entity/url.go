package entity

type URL struct {
	Key         string `json:"key" redis:"key"`
	OriginalURL string `json:"original_url" redis:"original_url"`
	Expires     string `json:"expires" redis:"expires"`
	Visits      int    `json:"visits" redis:"visits"`
}
