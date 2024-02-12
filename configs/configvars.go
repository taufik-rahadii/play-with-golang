package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfigVar(key string) string {
	value := os.Getenv(key)

	return value
}

func InitEnv(envPath string) error {
	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatal("ERROR LOADING ENV FILE")
	}

	return err
}
