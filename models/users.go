package models

import "time"

type User struct {
	ID           int        `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	ResetCode    *int       `json:"reset_code,omitempty"`
	FIO          *string    `json:"fio,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	LastLogin    *time.Time `json:"last_login,omitempty"`
	AvatarURL    *string    `json:"avatar_url,omitempty"`
	IsOnline     bool       `json:"is_online"`
	Bio          *string    `json:"bio,omitempty"`
}
