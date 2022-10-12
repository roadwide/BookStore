package configs

import "os"

var (
	JWTSecretKey string
	SQLiteFile   string
)

func init() {
	SQLiteFile = os.Getenv("SQLITE_FILE")
	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
}
