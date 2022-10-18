package queries

import (
	"backend/ent"
	"backend/ent/book"
	"context"
	"errors"
	"strings"
)

func (db *Query) CreateBook(userID, bookName, picURL string, bookPrice float64) (*ent.Book, error) {
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

// 根据book_id在查询数据库中查询对应的图片文件名
func (db *Query) GetBookImgFileNameById(bookID int) string {
	PicURL, err := db.Book.Query().Where(book.ID(bookID)).Select(book.FieldPicURL).String(context.Background())
	if err != nil {
		return ""
	}
	arr := strings.Split(PicURL, "/") // 以 / 分割返回数组，数组最后一个元素就是文件名
	return arr[len(arr)-1]
}

func (db *Query) DeleteBook(bookID int, userID string) (string, error) {
	// 判断用户名下是否有这本书，通过计数来实现，有就是1，没有就是0
	book_count, err := db.Book.Query().Where(book.ID(bookID)).Where(book.UserID(userID)).Count(context.Background())
	if err != nil {
		return "", err
	}
	if book_count == 1 {
		ImgFileName := db.GetBookImgFileNameById(bookID)
		return ImgFileName, db.Book.DeleteOneID(bookID).Exec(context.Background())
	} else {
		return "", errors.New("当前用户不存在这本书")
	}
}
