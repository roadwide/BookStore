package models

type AddBookRequest struct {
	UserID string  `json:"user_id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	PicURL string  `json:"pic_url"`
}

type GetBookRequest struct {
	UserID string `json:"user_id"`
}

type DeleteBookRequest struct {
	BookID int `json:"book_id"`
}
