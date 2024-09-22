package commentRepository

import (
	"errors"
	"log"

	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

// GetRepliesByMalID implements CommentRepository.
func (p *postgresRepository) GetRepliesByMalID(mal_id string) ([]entities.Comments, error) {
	replies := []entities.Comments{}
	if err := p.db.Where(&entities.Comments{MalID: mal_id}).
		Not("parent_id = ?", 0).Preload("User").
		Find(&replies).Error; err != nil {
		log.Printf("error: GetRepliesByMalID: %s\n", err.Error())
		return []entities.Comments{}, errors.New("error: find reply failed")
	}
	return replies, nil
}

// InsertOneComment implements CommentRepository.
func (p *postgresRepository) InsertOneComment(in *entities.Comments) error {
	if err := p.db.Create(&in).Error; err != nil {
		log.Printf("error: InsertOneComment: %s", err.Error())
		return errors.New("error: insert comment failed")
	}
	return nil
}

// GetCommentsByMal_id implements CommentRepository.
func (p *postgresRepository) GetParentCommentsByMalID(mal_id string) ([]entities.Comments, error) {
	parentComments := []entities.Comments{}
	if err := p.db.
		Where(map[string]interface{}{"mal_id": mal_id, "parent_id": 0}).
		Preload("User").
		Find(&parentComments).Error; err != nil {
		log.Printf("error: GetCommentAndReplyByMalID: %s\n", err.Error())
		return []entities.Comments{}, errors.New("error: find comment failed")
	}
	return parentComments, nil
}

func NewPostgresRepository(db *gorm.DB) CommentRepository {
	return &postgresRepository{db: db}
}
