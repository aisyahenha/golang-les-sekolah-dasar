package model

type User struct {
	BaseModel
	Username   string `json:"username"`
	Password   string `json:"password,omitempty"`
	Role       string `json:"role"`
	ResetToken string `json:"resetToken,omitempty"`
}
