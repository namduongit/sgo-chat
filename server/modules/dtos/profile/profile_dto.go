package profile

type ProfileRes struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Fullname string `json:"fullname"`
	Bio      string `json:"bio"`
	Gender   string `json:"gender"`
}
