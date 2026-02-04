package service

import repository "github/sgo-chat/modules/repositories"

type AuthService struct {
	accountRepo *repository.AccountRepository
	jwtSecret   string
}

type AuthServiceConfig struct {
	JWTSecret string
}

func NewAuthService(accountRepo *repository.AccountRepository, jwtSecret string) *AuthService {
	return &AuthService{
		accountRepo: accountRepo,
		jwtSecret:   jwtSecret,
	}
}
