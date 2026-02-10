package configs

import (
	"log"
	"os"
	"strconv"

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

	Crypto Bcrypt
}

func Load() *Variable {
	_ = godotenv.Load()

	cost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		log.Fatalln("[Variable] Cost value require type is int")
	}

	return &Variable{
		Port:      os.Getenv("PORT"),
		MongoURL:  os.Getenv("MONGO_URL"),
		DBname:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),

		Crypto: Bcrypt{
			Cost: cost,
		},
	}
}
