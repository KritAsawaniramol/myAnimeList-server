package userUsecase

import "github.com/kritAsawaniramol/myAnimeList-server/module/user"

type (
	UserUsecase interface {
		GetUserProfile(userID uint) (*user.UserProfileRes, error)
	}
)
