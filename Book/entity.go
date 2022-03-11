package Book

import "time"

type Book struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price int64 `json:"price"`
	Rating int64 `json:"rating"`
	CreatedAt time.Time
	UpdatedAt time.Time
}