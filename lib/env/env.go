package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Setup() {
	if err := godotenv.Load(); err != nil {
		log.Panicf("Error loading env file: %s", err.Error())
	}
}

func Get(key string) *parser {
	val := os.Getenv(key)
	return parse(val)
}
