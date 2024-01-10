package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Port               string
	LibraryDBHost      string
	LibraryDBPort      string
	CentralLibraryPort string
	CentralLibraryHost string
	City               string
}

func NewConfig() *Config {
	LoadEnv()

	return &Config{
		Port:               os.Getenv("LIBRARY_PORT"),
		LibraryDBHost:      os.Getenv("LIBRARY_DB_HOST"),
		LibraryDBPort:      os.Getenv("LIBRARY_DB_PORT"),
		CentralLibraryHost: os.Getenv("CENTRAL_LIBRARY_HOST"),
		CentralLibraryPort: os.Getenv("CENTRAL_LIBRARY_PORT"),
		City:               os.Getenv("CITY"),
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
