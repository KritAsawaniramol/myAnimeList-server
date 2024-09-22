package auth

import (
	"time"
)

type LoginWithOauth struct {
	Email              string `json:"email"`
	Name               string `json:"name"`
	AuthProviderName   string `json:"auth_provider"`
	AuthProviderUserID string `json:"auth_provider_user_id"`
	AvatarURL          string `json:"avatar_url"`
}

type LoginWithOauthRes struct {
	LoginWithOauth
	Credential *CredentialRes `json:"credential"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" validate:"required,jwt"`
	CredentialID uint   `json:"credential_id" validate:"required,gt=0"`
}

type CredentialRes struct {
	Id           uint      `json:"credential_id"`
	UserID       uint      `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
