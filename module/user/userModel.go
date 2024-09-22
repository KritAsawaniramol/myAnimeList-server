package user

type (
	UserProfileRes struct {
		Name         string `json:"name"`
		Email        string `json:"email"`
		AvatarURL    string `json:"avatar_url"`
	}
)
