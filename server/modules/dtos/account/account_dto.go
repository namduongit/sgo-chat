package account

import "github/sgo-chat/modules/dtos/profile"

type AccountRes struct {
	ID       string             `json:"id"`
	Email    string             `json:"email"`
	Profile  profile.ProfileRes `json:"profile"`
	CreateAt string             `json:"create_at"`
}

type CreateAccountReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
