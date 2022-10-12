package utils

import (
	"fmt"
)

func GetSQLiteURL(filename string) string {
	return fmt.Sprintf("file:%s?cache=shared&_fk=1", filename)
}
