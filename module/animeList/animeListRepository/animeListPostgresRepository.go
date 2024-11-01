package animeListRepository

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"github.com/kritAsawaniramol/myAnimeList-server/util"
	"gorm.io/gorm"
)

type animeListPostgresRepository struct {
	db *gorm.DB
}

// DeleteOneAnimeList implements AnimeListRepository.
func (a *animeListPostgresRepository) DeleteOneAnimeList(malID string, userID uint) error {
	animeList := &entities.AnimeLists{MalID: malID, UserID: userID}
	err := a.db.Unscoped().Delete(&animeList).Error
	fmt.Printf("animeList: %v\n", animeList)
	if err != nil {
		log.Printf("error: DeleteOneAnimeList: %s\n", err.Error())
		return errors.New("error: remove anime from my list failed")
	}
	return nil
}

// UpdateOneAnimeList implements AnimeListRepository.
func (a *animeListPostgresRepository) UpdateOneAnimeList(malID string, userID uint, in *entities.AnimeLists) error {
	condition := &entities.AnimeLists{}
	condition.MalID = malID
	condition.UserID = userID
	if err := a.db.Model(&condition).Where(&condition).Select("Status", "EpisodesCount", "Score").Updates(in).Error; err != nil {
		log.Printf("error: UpdateOneAnimeList: %s\n", err.Error())
		return errors.New("error: update anime list failed")
	}
	return nil
}

// GetAnimeListByUserID implements AnimeListRepository.
func (a *animeListPostgresRepository) GetAnimeListByUserID(userID uint) ([]entities.AnimeLists, error) {
	animeList := []entities.AnimeLists{}
	if err := a.db.Where(&entities.AnimeLists{UserID: userID}).Find(&animeList).Error; err != nil {
		log.Printf("error: GetAnimeListByUserID: %s\n", err.Error())
		return []entities.AnimeLists{}, errors.New("error: anime list not found")
	}
	return animeList, nil
}

// GetOneAnimeList implements AnimeListRepository.
func (a *animeListPostgresRepository) GetOneAnimeList(in *entities.AnimeLists) (*entities.AnimeLists, error) {
	animeList := &entities.AnimeLists{}
	if err := a.db.Where(&in).First(&animeList).Error; err != nil {
		log.Printf("error: GetOneAnimeList: %s\n", err.Error())
		return &entities.AnimeLists{}, errors.New("error: not found in anime list")
	}
	return animeList, nil
}

// InsertOneAnimeToAnimeList implements AnimeListRepository.
func (a *animeListPostgresRepository) InsertOneAnimeToAnimeList(in *entities.AnimeLists) error {
	if err := a.db.Create(&in).Error; err != nil {
		log.Printf("error: InsertOneAnimeToAnimeList: %s\n", err.Error())
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errors.New("error: this anime is already in list")
		}
		return errors.New("error: Add anime to my list failed")
	}
	util.PrintObjInJson(in)
	return nil
}

func NewAnimeRepository(db *gorm.DB) AnimeListRepository {
	return &animeListPostgresRepository{db: db}
}
