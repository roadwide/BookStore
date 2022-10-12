package queries

import (
	"backend/ent"
	"backend/ent/book"
	"context"
)

func (db *Query) CrateBook(userID, bookName, picURL string, bookPrice float64) (*ent.Book, error) {
	return db.Book.Create().
		SetName(bookName).
		SetUserID(userID).
		SetPrice(bookPrice).
		SetPicURL(picURL).
		Save(context.Background())
}

func (db *Query) GetBook(userID string) ([]*ent.Book, error) {
	return db.Book.Query().
		Where(book.UserID(userID)).
		All(context.Background())
}

func (db *Query) GetAllBook() ([]*ent.Book, error) {
	return db.Book.Query().All(context.Background())
}
