package middlewareRepository

import (
	"errors"
	"log"

	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"gorm.io/gorm"
)

type middlewareRepository struct {
	db *gorm.DB
}

// UpdateOneUserCredential implements MidlerwareRepository.
func (m *middlewareRepository) UpdateOneUserCredential(id uint, in *entities.Credential) error {
	model := &entities.Credential{}
	model.ID = id
	if err := m.db.Model(&model).Updates(in).Error; err != nil {
		log.Printf("error: UpdateOneUserCredential: %s\n", err.Error())
		return errors.New("error: update credential failed")
	}
	return nil
}

// AccessTokenSearch implements MidlerwareRepository.
func (m *middlewareRepository) AccessTokenSearch(accessToken string) (*entities.Credential, error) {
	credential := &entities.Credential{}
	if err := m.db.Where(&entities.Credential{AccessToken: accessToken}).Last(credential).Error; err != nil {
		log.Printf("error: AccessTokenSearch: %s\n", err.Error())
		return nil, errors.New("error: credential not found")
	}
	return credential, nil
}

func (m *middlewareRepository) GetOneUserByID(userID uint) (*entities.User, error) {
	user := &entities.User{}
	user.ID = userID
	if err := m.db.Where(user).First(&user).Error; err != nil {
		log.Printf("error: GetOneUser: %s\n", err.Error())
		return nil, errors.New("error: user not found")
	}
	return user, nil
}

// GetOneUserCredentialByID implements AuthServiceRepository.
func (m *middlewareRepository) GetOneUserCredentialByID(id uint) (*entities.Credential, error) {
	credential := &entities.Credential{}
	if err := m.db.Find(&credential, "id = ?", id).Error; err != nil {
		log.Printf("error: GetOneUserCredentialByID: %s\n", err.Error())
		return nil, errors.New("error: user's credentail not found")
	}
	return credential, nil
}

func NewMiddlewareRepository(db *gorm.DB) MidlerwareRepository {
	return &middlewareRepository{db: db}
}
