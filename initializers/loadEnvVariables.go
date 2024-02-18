package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
)

func LoadEnvVariables() {
	// Load .env file
	pathDir, err := filepath.Abs(filepath.Dir("."))
	err = godotenv.Load(filepath.Join(pathDir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
