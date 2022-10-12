package queries

import (
	"backend/ent"
	"backend/pkg/configs"
	"backend/pkg/utils"
	"backend/platform"
)

type Query struct {
	*ent.Client
}

var (
	DataBase *Query
)

func init() {
	DataBase = &Query{platform.GetDBClient(utils.GetSQLiteURL(configs.SQLiteFile))}
}
