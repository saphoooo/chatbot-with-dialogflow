package views

// WikiSummary ...
type WikiSummary struct {
	Thumbnail Thumbnail `json:"thumbnail,omitempty"`
	Extract   string    `json:"extract,omitempty"`
}

// Thumbnail ...
type Thumbnail struct {
	Source string `json:"source,omitempty"`
}
