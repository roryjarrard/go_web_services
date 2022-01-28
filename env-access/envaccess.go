package envaccess

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	env := GetKey("env")
	err := godotenv.Load(".env." + env)
	return err
}

func GetKey(key string) string {
	return os.Getenv(key)
}
