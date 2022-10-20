package models

type AddBookRequest struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	PicURL string  `json:"pic_url"`
	Token  string  `json:"token"`
}

type GetBookRequest struct {
	UserID string `json:"user_id"`
}

type DeleteBookRequest struct {
	BookID int    `json:"book_id"`
	Token  string `json:"token"`
}
