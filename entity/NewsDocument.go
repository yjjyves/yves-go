package entity

type NewsDocument struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Category    string   `json:"category"`
	PublishedAt string   `json:"published_at"`
	Tags        []string `json:"tags"`
	Speech      string   `json:"speech"`
	Score       float64  `json:"score"`
}
