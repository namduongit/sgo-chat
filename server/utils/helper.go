package utils

import (
	"strconv"

	"github/sgo-chat/internals/configs"
)

type HelperUtils struct {
	Bcrypt *Bcrypt
	Jsonwt *JsonWebToken
}

func NewHelperUtils(cfg *configs.Variable) *HelperUtils {
	cost, err := strconv.Atoi(cfg.Cost)
	if err != nil {
		// Define a default cost value if conversion fails
		cost = 10
	}
	jwtExpire, err := strconv.ParseInt(cfg.JWTExpire, 10, 64)
	if err != nil {
		// Define a default JWT expire value if conversion fails
		jwtExpire = 3600
	}
	return &HelperUtils{
		Bcrypt: &Bcrypt{Cost: cost},
		Jsonwt: &JsonWebToken{
			JWTSecret: cfg.JWTSecret,
			JWTExpire: jwtExpire,
		},
	}
}
