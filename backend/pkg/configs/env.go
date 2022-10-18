package configs

var (
	JWTSecretKey string
	SQLiteFile   string
)

func init() {
	// SQLiteFile = os.Getenv("SQLITE_FILE")
	// JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	SQLiteFile = "data.s3db"
	JWTSecretKey = "MxDNSc0AnSSoU7WUUWh9i"
}
