package comment

import "time"

type (
	PostCommentReq struct {
		UserID   uint
		MalID    string `json:"mal_id" validate:"required,min=1"`
		Content  string `json:"content" varlidate:"required,max=3000"`
		ParentID uint   `json:"parent_id" varlidate:"gt=0"`
	}

	GetAnimeCommentsReq struct {
		MalID string `json:"mal_id" validate:"required,min=1"`
	}

	GetAnimeCommentsRes struct {
		Comments []Comment `json:"comments"`
	}

	Comment struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		AvatarUrl string    `json:"avatar_url"`
		Content   string    `json:"content"`
		Timestamp time.Time `json:"timestamp"`
		Replies   []Reply   `json:"replies"`
	}

	Reply struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		AvatarUrl string    `json:"avatar_url"`
		Content   string    `json:"content"`
		Timestamp time.Time `json:"timestamp"`
	}
)
