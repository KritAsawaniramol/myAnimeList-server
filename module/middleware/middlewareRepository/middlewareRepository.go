package middlewareRepository

import "github.com/kritAsawaniramol/myAnimeList-server/entities"

type MidlerwareRepository interface {
	AccessTokenSearch(accessToken string) (*entities.Credential, error)
	GetOneUserByID(userID uint) (*entities.User, error)
	UpdateOneUserCredential(id uint, in *entities.Credential) error
	GetOneUserCredentialByID(id uint) (*entities.Credential, error)
}
