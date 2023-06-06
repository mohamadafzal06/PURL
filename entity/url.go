package entity

type URL struct {
	Key         string `json:"key" redis:"key"`
	OriginalURL string `json:"original_url" redis:"original_url"`
	// TODO: the type of Expires can be time.Time
	Expires string `json:"expires" redis:"expires"`
	Visits  int    `json:"visits" redis:"visits"`
}
