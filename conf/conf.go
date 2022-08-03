package conf

import (
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}
