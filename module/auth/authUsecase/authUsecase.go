package authUsecase

import (
	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth"
)

type (
	AuthUsecase interface {
		LoginWithOauth(cfg *config.Config, in *auth.LoginWithOauth) (*auth.LoginWithOauthRes, error)
		Logout(credentialID uint) error
		RefreshToken(cfg *config.Config, credentialID uint, refreshToken string) (*auth.CredentialRes, error)
	}
)
