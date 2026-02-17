package services

import (
	"github/sgo-chat/internals/configs/errors"
	"github/sgo-chat/models"
	"github/sgo-chat/modules/dtos/auth"
	"github/sgo-chat/modules/repositories"
	"github/sgo-chat/utils"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	accountRepo *repositories.AccountRepository
	helper      *utils.HelperUtils
}

func NewAuthService(accountRepo *repositories.AccountRepository, helper *utils.HelperUtils) *AuthService {
	return &AuthService{
		accountRepo: accountRepo,
		helper:      helper,
	}
}

func (as *AuthService) Register(c *gin.Context) (*auth.RegisterRes, error) {
	ctx := c.Request.Context()

	req := auth.RegisterAccountReq{}
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

	hashedPassword, err := as.helper.Bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	result, err = as.accountRepo.Create(ctx, &models.Account{
		Email:    req.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		return nil, err
	}

	res := &auth.RegisterRes{
		ID:       result.ID.Hex(),
		Email:    result.Email,
		CreateAt: result.CreateAt.String(),
	}

	return res, nil
}

func (as *AuthService) Login(c *gin.Context) (*auth.LoginRes, error) {
	ctx := c.Request.Context()

	req := auth.LoginAccountReq{}

	// Validate request body
	if err := utils.ShouldBindReq(c, &req); err != nil {
		return nil, err
	}

	account, err := as.accountRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.BadRequestError("Invalid email or password")
	}

	err = as.helper.Bcrypt.ComparePassword(account.Password, req.Password)
	if err != nil {
		return nil, errors.BadRequestError("Invalid email or password")
	}

	token, exp, iat, err := as.helper.Jsonwt.GenerateToken(map[string]string{
		"id":    account.ID.Hex(),
		"email": account.Email,
	})
	if err != nil {
		c.Error(err)
		return nil, err
	}

	res := &auth.LoginRes{
		ID:    account.ID.Hex(),
		Email: account.Email,
		Token: token,
		Time: auth.LifeTime{
			ExpiratedAt: exp,
			IssuedAt:    iat,
		},
	}

	return res, nil
}
