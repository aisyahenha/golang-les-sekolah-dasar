package model

import "time"

type User struct {
	BaseModel
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	ResetToken string `json:"resetToken,omitempty"`
}
type UserRespon struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	ResetToken string    `json:"resetToken,omitempty"`
}
