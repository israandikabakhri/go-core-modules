package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL string
}

func LoadConfig() *Config {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return &Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
    }
}