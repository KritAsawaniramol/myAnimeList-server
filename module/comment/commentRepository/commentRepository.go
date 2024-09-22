package commentRepository

import "github.com/kritAsawaniramol/myAnimeList-server/entities"

type CommentRepository interface {
	InsertOneComment(in *entities.Comments) error
	GetParentCommentsByMalID(mal_id string) ([]entities.Comments, error)
	GetRepliesByMalID(mal_id string) ([]entities.Comments, error)
}
