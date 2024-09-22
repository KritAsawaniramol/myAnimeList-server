package authRepository

import (
	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/jwtAuth"
)

type (
	AuthServiceRepository interface {
		GetOneUserByAuthProvider(authProviderName string, authProviderUserID string) (*entities.User, error)
		AccessToken(cfg *config.Config, claims *jwtAuth.Claims) string
		RefreshToken(cfg *config.Config, claims *jwtAuth.Claims) string
		CreateOneUser(in *entities.User) (uint, error)
		InserOneUserCredential(in *entities.Credential) (uint, error)
		GetOneUserCredentialByID(id uint) (*entities.Credential, error)
		DeleteOneUserCredentialByID(id uint) error
		UpdateOneUserCredential(condition *entities.Credential, in *entities.Credential) error
		GetOneUserByID(userID uint) (*entities.User, error)
		UpdateOneUser(in *entities.User) error
	}
)
