package animeListUsecase

import (
	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"github.com/kritAsawaniramol/myAnimeList-server/module/animeList"
	"github.com/kritAsawaniramol/myAnimeList-server/module/animeList/animeListRepository"
)

type animeListUsecaseImpl struct {
	animeListRepository animeListRepository.AnimeListRepository
}

// RemoveOneAnimeInList implements AnimeListUsecase.
func (a *animeListUsecaseImpl) RemoveOneAnimeInList(malID string, userID uint) error {
	if err := a.animeListRepository.DeleteOneAnimeList(malID, userID); err != nil {
		return err
	}
	return nil
}

// GetAnimeList implements AnimeListUsecase.
func (a *animeListUsecaseImpl) GetAnimeList(userID uint) (*animeList.GetAnimeListRes, error) {
	list, err := a.animeListRepository.GetAnimeListByUserID(userID)
	if err != nil {
		return &animeList.GetAnimeListRes{AnimeList: []animeList.AnimeList{}}, err
	}
	res := []animeList.AnimeList{}
	for _, v := range list {
		res = append(res, animeList.AnimeList{
			MalID:         v.MalID,
			Status:        v.Status,
			EpisodesCount: v.EpisodesCount,
			Score:         v.Score,
		})
	}
	return &animeList.GetAnimeListRes{AnimeList: res}, nil
}

// UpdateOneAnimeList implements AnimeListUsecase.
func (a *animeListUsecaseImpl) UpdateOneAnimeList(req *animeList.UpdateOneAnimeListReq) (*animeList.UpdateOneAnimeListRes, error) {
	record, err := a.animeListRepository.GetOneAnimeList(&entities.AnimeLists{
		MalID:  req.MalID,
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	if req.EpisodesCount != nil {
		record.EpisodesCount = *req.EpisodesCount
	}
	if req.Score != nil {
		record.Score = *req.Score
	}
	if req.Status != nil {
		record.Status = *req.Status
	}

	if err := a.animeListRepository.UpdateOneAnimeList(req.MalID, req.UserID, record); err != nil {
		return nil, err
	}

	newRecord, err := a.animeListRepository.GetOneAnimeList(&entities.AnimeLists{
		MalID:  req.MalID,
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &animeList.UpdateOneAnimeListRes{
		UserID:        newRecord.UserID,
		MalID:         newRecord.MalID,
		Status:        newRecord.Status,
		EpisodesCount: newRecord.EpisodesCount,
		Score:         newRecord.Score,
	}, nil
}

// GetOneAnimeList implements AnimeListUsecase.
func (a *animeListUsecaseImpl) GetOneAnimeList(malID string, userID uint) (*animeList.GetOneAnimeListRes, error) {
	record, err := a.animeListRepository.GetOneAnimeList(&entities.AnimeLists{
		MalID:  malID,
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	res := &animeList.GetOneAnimeListRes{
		UserID:        record.UserID,
		MalID:         record.MalID,
		Status:        record.Status,
		EpisodesCount: record.EpisodesCount,
		Score:         record.Score,
	}
	return res, nil
}

// addAnimeToMyList implements AnimeListUsecase.
func (a *animeListUsecaseImpl) AddAnimeToMyList(req *animeList.AddAnimeToMyListReq) error {
	if err := a.animeListRepository.InsertOneAnimeToAnimeList(&entities.AnimeLists{
		MalID:         req.MalID,
		UserID:        req.UserID,
		Status:        "plan-to-watch",
		Score:         0,
		EpisodesCount: 0,
	}); err != nil {
		return err
	} else {
		return nil
	}
}

func NewAnimeListUsecase(
	animeListRepository animeListRepository.AnimeListRepository) AnimeListUsecase {
	return &animeListUsecaseImpl{
		animeListRepository: animeListRepository,
	}
}
