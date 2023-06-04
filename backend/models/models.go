package models

type Repository struct {
	Name        string `json:"Name"`
	Languange   string `json:"Language"`
	Stars       int    `json:"Stars"`
	Description string `json:"Description"`
	Link        string `json:"Link"`
	NumOfLikes  int    `json:"NumOfLikes"`
}

type Database []Repository
