package commentUsecase

import "github.com/kritAsawaniramol/myAnimeList-server/module/comment"

type CommentUsecase interface {
	PostComment(in *comment.PostCommentReq) error
	GetCommentsByMalID(malID string) (*comment.GetAnimeCommentsRes, error)
}