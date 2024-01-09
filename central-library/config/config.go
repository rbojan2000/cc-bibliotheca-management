package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Port                 string
	CentralLibraryDBHost string
	CentralLibraryDBPort string
}

func NewConfig() *Config {
	LoadEnv()

	return &Config{
		Port:                 os.Getenv("CENTRAL_LIBRARY_PORT"),
		CentralLibraryDBHost: os.Getenv("CENTRAL_LIBRARY_DB_HOST"),
		CentralLibraryDBPort: os.Getenv("CENTRAL_LIBRARY_DB_PORT"),
	}
}

func LoadEnv() {
	var envPath string

	envPath = filepath.FromSlash("./.env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}
