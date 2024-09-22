package userRepository

import "github.com/kritAsawaniramol/myAnimeList-server/entities"

type UserRepository interface {
	GetOneUser(userID uint) (*entities.User, error)
}