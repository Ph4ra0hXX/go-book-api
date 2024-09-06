package model

type Book struct {
	ID     int    `json:"id"`
	Image  string `json:"image"`
	Author string `json:"author"`
	Title  string `json:"title"`
}
