package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Bcrypt struct {
	Cost int
}

type Variable struct {
	Port string

	MongoURL  string
	DBname    string
	JWTSecret string
	JWTExpire string
	Cost      string
}

func Load() *Variable {
	_ = godotenv.Load()
	return &Variable{
		Port:      os.Getenv("PORT"),
		MongoURL:  os.Getenv("MONGO_URL"),
		DBname:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpire: os.Getenv("JWT_EXPIRE"),
		Cost:      os.Getenv("COST"),
	}
}
