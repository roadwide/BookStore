package platform

import (
	"backend/ent"
	"context"
	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func GetDBClient(url string) *ent.Client {
	client, err := ent.Open(dialect.SQLite, url)
	if err != nil {
		panic(err.Error())
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return client
}
