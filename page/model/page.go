package model

type Page struct {
	BookID     int    `json:"book_id"`
	PageNumber int    `json:"page_number"`
	Text       string `json:"text"`
}
