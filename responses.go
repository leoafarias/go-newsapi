package newsapi

import "time"

type article struct {
	Source      source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
}

type source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Category    string `json:"category,omitempty"`
	Language    string `json:"language,omitempty"`
	Country     string `json:"country,omitempty"`
}

// ArticlesResponse is a response expected on endpoint that return articles
type ArticlesResponse struct {
	Status       string `json:"status"`
	Code         string `json:"code"`
	Message      string `json:"message"`
	TotalResults int    `json:"totalResults"`
	Articles     []article
}

// SourcesResponse is a response expected on endpoints that return sources
type SourcesResponse struct {
	Status       string `json:"status"`
	Code         string `json:"code"`
	Message      string `json:"message"`
	TotalResults int    `json:"totalResults"`
	Sources      []source
}
