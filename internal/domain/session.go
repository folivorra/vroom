package domain

import "time"

type Session struct {
	Id           int64     `json:"id"`
	UserId       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
	Revoked      bool      `json:"revoked"`
}
