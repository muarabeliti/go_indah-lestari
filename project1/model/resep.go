package models

type Resep struct {
	Id          uint     `json:"id"`
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}
