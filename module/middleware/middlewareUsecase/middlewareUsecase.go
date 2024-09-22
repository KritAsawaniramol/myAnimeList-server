package middlewareUsecase

import (
	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth"
)

type MiddlewareUsecaseService interface {
	JwtAuthorization(cfg *config.Config, accessToken string) (uint, error)
	RefreshToken(cfg *config.Config, credentialID uint,refreshToken string) (*auth.CredentialRes, error) 
}
