package models

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}
