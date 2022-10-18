package queries

import (
	"testing"
)

func TestDeleteBook(t *testing.T) {
	DataBase.DeleteBook(2, "1234")
}
