package animeListRepository

import "github.com/kritAsawaniramol/myAnimeList-server/entities"

type AnimeListRepository interface {
	InsertOneAnimeToAnimeList(in *entities.AnimeLists) error
	GetAnimeListByUserID(userID uint) ([]entities.AnimeLists, error)
	GetOneAnimeList(in *entities.AnimeLists) (*entities.AnimeLists, error)
	UpdateOneAnimeList(malID string, userID uint, in *entities.AnimeLists) error
	DeleteOneAnimeList(malID string, userID uint) error
}
