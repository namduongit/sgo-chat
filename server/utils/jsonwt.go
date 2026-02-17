package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JsonWebToken struct {
	JWTSecret string
	JWTExpire int64
}

type claims struct {
	jwt.RegisteredClaims
	Subject any `json:"subject"`
}

func (jswt *JsonWebToken) GenerateToken(subject any) (string, int64, int64, error) {
	claims := &claims{
		Subject: subject,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jswt.JWTExpire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "sgo-chat",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte(jswt.JWTSecret))
	if err != nil {
		return "", 0, 0, err
	}

	return signedToken, claims.ExpiresAt.Unix(), time.Now().Unix(), nil
}
