package userRepository

import (
	"errors"
	"log"

	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"gorm.io/gorm"
)

type userPostgrestRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userPostgrestRepository{db: db}
}

// GetOneUser implements UserRepository.
func (u *userPostgrestRepository) GetOneUser(userID uint) (*entities.User, error) {
	user := &entities.User{}
	user.ID = userID
	if err := u.db.Where(user).First(&user).Error; err != nil {
		log.Printf("error: GetOneUser: %s\n", err.Error())
		return nil, errors.New("error: user not found")
	}
	return user, nil
}


