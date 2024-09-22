package animeList

type (
	AddAnimeToMyListReq struct {
		UserID uint
		MalID  string `json:"mal_id" validate:"required,min=1"`
	}

	GetOneAnimeListRes struct {
		UserID        uint   `json:"user_id"`
		MalID         string `json:"mal_id"`
		Status        string `json:"status"`
		EpisodesCount uint   `json:"episodes_count"`
		Score         int    `json:"score"`
	}

	UpdateOneAnimeListReq struct {
		UserID        uint
		MalID         string
		Status        *string `json:"status,omitempty" validate:"omitempty,is_myanime_status"`
		EpisodesCount *uint   `json:"episodes_count,omitempty" validate:"omitempty,gte=0"`
		Score         *int    `json:"score,omitempty" validate:"omitempty,gte=0,lte=10"`
	}

	UpdateOneAnimeListRes struct {
		UserID        uint   `json:"user_id"`
		MalID         string `json:"mal_id"`
		Status        string `json:"status"`
		EpisodesCount uint   `json:"episodes_count"`
		Score         int    `json:"score"`
	}

	AnimeList struct {
		MalID         string `json:"mal_id"`
		Status        string `json:"status"`
		EpisodesCount uint   `json:"episodes_count"`
		Score         int    `json:"score"`
	}

	GetAnimeListRes struct {
		AnimeList []AnimeList `json:"anime_list"`
	}
)
