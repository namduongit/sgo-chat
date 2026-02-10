package services

import (
	"github/sgo-chat/internals/configs/errors"
	"github/sgo-chat/modules/dtos/account"
	"github/sgo-chat/modules/repositories"
	"github/sgo-chat/utils"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	accountRepo *repositories.AccountRepository
	jwtSecret   string
}

func NewAuthService(accountRepo *repositories.AccountRepository, jwtSecret string) *AuthService {
	return &AuthService{
		accountRepo: accountRepo,
		jwtSecret:   jwtSecret,
	}
}

func (as *AuthService) Register(c *gin.Context) (*account.AccountRes, error) {
	ctx := c.Request.Context()

	req := account.CreateAccountReq{}
	// Validate request body
	if err := utils.ShouldBindReq(c, &req); err != nil {
		return nil, err
	}
	// Check if email already exists
	result, err := as.accountRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if result != nil {
		return nil, errors.BadRequestError("Email already exists")
	}

	result, err = as.accountRepo.Create(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	res := &account.AccountRes{
		ID:       result.ID.Hex(),
		Email:    result.Email,
		CreateAt: result.CreateAt.String(),
	}

	return res, nil
}
