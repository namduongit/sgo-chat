package auth

type LifeTime struct {
	ExpiratedAt int64 `json:"expired_at"`
	IssuedAt    int64 `json:"issued_at"`
}

type LoginRes struct {
	ID    string   `json:"id"`
	Email string   `json:"email"`
	Token string   `json:"token"`
	Time  LifeTime `json:"time"`
}

type RegisterRes struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	CreateAt string `json:"create_at"`
}

type RegisterAccountReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginAccountReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
