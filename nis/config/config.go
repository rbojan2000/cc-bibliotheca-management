package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Port               string
	NisLibraryDBHost   string
	NisLibraryDBPort   string
	CentralLibraryPort string
	CentralLibraryHost string
}

func NewConfig() *Config {
	LoadEnv()

	return &Config{
		Port:               os.Getenv("NIS_LIBRARY_PORT"),
		NisLibraryDBHost:   os.Getenv("NIS_LIBRARY_DB_HOST"),
		NisLibraryDBPort:   os.Getenv("NIS_LIBRARY_DB_PORT"),
		CentralLibraryHost: os.Getenv("CENTRAL_LIBRARY_HOST"),
		CentralLibraryPort: os.Getenv("CENTRAL_LIBRARY_PORT"),
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
