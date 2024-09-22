package animeListUsecase

import "github.com/kritAsawaniramol/myAnimeList-server/module/animeList"

type AnimeListUsecase interface {
	AddAnimeToMyList(req *animeList.AddAnimeToMyListReq) error
	GetOneAnimeList(malID string, userID uint) (*animeList.GetOneAnimeListRes, error)
	UpdateOneAnimeList(req *animeList.UpdateOneAnimeListReq) (*animeList.UpdateOneAnimeListRes, error)
	GetAnimeList(userID uint) (*animeList.GetAnimeListRes, error)
	RemoveOneAnimeInList(malID string, userID uint) error
}
