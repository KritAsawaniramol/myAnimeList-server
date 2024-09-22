package commentUsecase

import (
	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"github.com/kritAsawaniramol/myAnimeList-server/module/comment"
	"github.com/kritAsawaniramol/myAnimeList-server/module/comment/commentRepository"
)

type commentUsecaseImpl struct {
	commentRepository commentRepository.CommentRepository
}

// GetCommentsByMalID implements CommentUsecase.
func (c *commentUsecaseImpl) GetCommentsByMalID(malID string) (*comment.GetAnimeCommentsRes, error) {
	parentComments, err := c.commentRepository.GetParentCommentsByMalID(malID)
	if err != nil {
		return &comment.GetAnimeCommentsRes{Comments: []comment.Comment{}}, err
	}

	replies, err := c.commentRepository.GetRepliesByMalID(malID)
	if err != nil {
		return &comment.GetAnimeCommentsRes{Comments: []comment.Comment{}}, err
	}

	m := map[uint][]comment.Reply{}
	for _, r := range replies {
		_, ok := m[r.ParentID]
		if ok {
			m[r.ParentID] = append(m[r.ParentID], comment.Reply{
				ID:        r.ID,
				Name:      r.User.Name,
				AvatarUrl: r.User.AvatarURL,
				Content:   r.Content,
				Timestamp: r.CreatedAt,
			})
		} else {
			m[r.ParentID] = []comment.Reply{
				{
					ID:        r.ID,
					Name:      r.User.Name,
					AvatarUrl: r.User.AvatarURL,
					Content:   r.Content,
					Timestamp: r.CreatedAt,
				}}
		}
	}

	res := comment.GetAnimeCommentsRes{Comments: []comment.Comment{}}

	for _, p := range parentComments {
		c := comment.Comment{
			ID:        p.ID,
			Name:      p.User.Name,
			AvatarUrl: p.User.AvatarURL,
			Content:   p.Content,
			Timestamp: p.CreatedAt,
		}
		if r, ok := m[p.ID]; ok {
			c.Replies = r
		} else {
			c.Replies = []comment.Reply{}
		}
		res.Comments = append(res.Comments, c)
	}
	return &res, nil
}

func NewCommentUsecase(
	commentRepository commentRepository.CommentRepository,
) CommentUsecase {
	return &commentUsecaseImpl{
		commentRepository: commentRepository,
	}
}

// postComment implements CommentUsecase.
func (c *commentUsecaseImpl) PostComment(in *comment.PostCommentReq) error {
	if err := c.commentRepository.InsertOneComment(&entities.Comments{
		UserID:   in.UserID,
		MalID:    in.MalID,
		Content:  in.Content,
		ParentID: in.ParentID,
	}); err != nil {
		return err
	}
	return nil
}
