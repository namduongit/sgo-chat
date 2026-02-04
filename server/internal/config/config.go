package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	MongoURL  string
	DBname    string
	JWTSecret string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:      os.Getenv("PORT"),
		MongoURL:  os.Getenv("MONGO_URL"),
		DBname:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
