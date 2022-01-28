package envaccess

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	env := GetKey("env")
	return godotenv.Load(".env." + env)
}

func GetKey(key string) string {
	return os.Getenv(key)
}
