package middlewareUsecase

import (
	"log"

	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth"
	"github.com/kritAsawaniramol/myAnimeList-server/module/middleware/middlewareRepository"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/jwtAuth"
)

type middlerwareUsecase struct {
	middlewareRepository middlewareRepository.MidlerwareRepository
}

// RefreshToken implements MiddlewareUsecaseService.
func (m *middlerwareUsecase) RefreshToken(cfg *config.Config, credentialID uint, refreshToken string) (*auth.CredentialRes, error) {
	log.Println("RefreshToken()")

	//PareseToken
	claims, err := jwtAuth.ParseToken(cfg.Jwt.RefreshSecretKey, refreshToken)
	if err != nil {
		log.Printf("error: RefreshToken: %s\n", err.Error())
		return nil, err
	}

	user, err := m.middlewareRepository.GetOneUserByID(claims.UserID)
	if err != nil {
		log.Printf("error: RefreshToken: %s\n", err.Error())
		return nil, err
	}

	accessToken := jwtAuth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtAuth.Claims{
		UserID: user.ID,
	}).SignToken()

	log.Println("Reload Token refreshToken: %s\n", refreshToken)
	refreshToken = jwtAuth.ReloadToken(cfg.Jwt.RefreshSecretKey, claims.ExpiresAt.Unix(), &jwtAuth.Claims{
		UserID: user.ID,
	})

	log.Println("Reload Token refreshToken: %s\n", refreshToken)

	if err := m.middlewareRepository.UpdateOneUserCredential(credentialID, &entities.Credential{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}); err != nil {
		return nil, err
	}

	credential, err := m.middlewareRepository.GetOneUserCredentialByID(credentialID)
	if err != nil {
		return nil, err
	}

	return &auth.CredentialRes{
		Id:           credential.ID,
		UserID:       credential.UserID,
		AccessToken:  credential.AccessToken,
		RefreshToken: credential.RefreshToken,
		CreatedAt:    credential.CreatedAt,
		UpdatedAt:    credential.UpdatedAt,
	}, nil
}

// JwtAuthorization implements MiddlewareUsecaseService.
func (m *middlerwareUsecase) JwtAuthorization(cfg *config.Config, accessToken string) (uint, error) {
	claims, err := jwtAuth.ParseToken(cfg.Jwt.AccessSecretKey, accessToken)
	if err != nil {
		return 0, err
	}

	if _, err := m.middlewareRepository.AccessTokenSearch(accessToken); err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MidlerwareRepository) MiddlewareUsecaseService {
	return &middlerwareUsecase{
		middlewareRepository: middlewareRepository,
	}
}
