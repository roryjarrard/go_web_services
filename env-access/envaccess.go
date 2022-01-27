package envaccess

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	env := os.Args[1]
	godotenv.Load(".env." + env)

}

func GetKey(key string) string {
	return os.Getenv(key)
}
