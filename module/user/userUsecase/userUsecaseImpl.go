package userUsecase

import (
	"github.com/kritAsawaniramol/myAnimeList-server/module/user"
	"github.com/kritAsawaniramol/myAnimeList-server/module/user/userRepository"
)

type userUsecaseImpl struct {
	userRepository userRepository.UserRepository
}

func NewUserUsecase(userRepository userRepository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepository: userRepository}
}

// GetUserProfile implements UserUsecase.
func (u *userUsecaseImpl) GetUserProfile(userID uint) (*user.UserProfileRes, error) {
	userRecord, err := u.userRepository.GetOneUser(userID)
	if err != nil {
		return nil, err
	}
	return &user.UserProfileRes{
		Name:      userRecord.Name,
		Email:     userRecord.Email,
		AvatarURL: userRecord.AvatarURL,
	}, nil
}
