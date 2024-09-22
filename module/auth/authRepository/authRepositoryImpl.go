package authRepository

import (
	"errors"
	"log"

	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/jwtAuth"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) GetOneUserByID(userID uint) (*entities.User, error) {
	user := &entities.User{}
	user.ID = userID
	if err := a.db.Where(user).First(&user).Error; err != nil {
		log.Printf("error: GetOneUser: %s\n", err.Error())
		return nil, errors.New("error: user not found")
	}
	return user, nil
}

// UpdateOneUserCredential implements MidlerwareRepository.
func (a *authRepository) UpdateOneUserCredential(condition *entities.Credential, in *entities.Credential) error {
	// model := &entities.Credential{}
	// model.ID = id
	if err := a.db.Model(&condition).Updates(in).Error; err != nil {
		log.Printf("error: UpdateOneUserCredential: %s\n", err.Error())
		return errors.New("error: update credential failed")
	}
	return nil
}

// DeleteOneUserCredentialByID implements AuthServiceRepository.
func (a *authRepository) DeleteOneUserCredentialByID(id uint) error {
	if err := a.db.Delete(&entities.Credential{}, id).Error; err != nil {
		log.Printf("error: DeleteOneUserCredentialByID: %s\n", err.Error())
		return errors.New("error: delete user's credential failed")
	}
	return nil
}

func NewAuthRepository(db *gorm.DB) AuthServiceRepository {
	return &authRepository{db: db}
}

// GetOneUserCredentialByID implements AuthServiceRepository.
func (a *authRepository) GetOneUserCredentialByID(id uint) (*entities.Credential, error) {
	credential := &entities.Credential{}
	if err := a.db.Find(&credential, "id = ?", id).Error; err != nil {
		log.Printf("error: GetOneUserCredentialByID: %s\n", err.Error())
		return nil, errors.New("error: user's credentail not found")
	}
	return credential, nil
}

// InserOneUserCredential implements AuthServiceRepository.
func (a *authRepository) InserOneUserCredential(in *entities.Credential) (uint, error) {
	if err := a.db.Create(in).Error; err != nil {
		return 0, errors.New("error: InserOneUserCredential: insert new user credential failed")
	}
	return in.ID, nil
}

// CreateOneUser implements AuthServiceRepository.
func (a *authRepository) CreateOneUser(in *entities.User) (uint, error) {
	if err := a.db.Create(in).Error; err != nil {
		return 0, errors.New("error: CreateOneUser: create new user failed")
	}
	return in.ID, nil
}

// AccessToken implements AuthRepository.
func (a *authRepository) AccessToken(cfg *config.Config, claims *jwtAuth.Claims) string {
	return jwtAuth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtAuth.Claims{
		UserID: claims.UserID,
	}).SignToken()
}

// RefreshToken implements AuthRepository.
func (a *authRepository) RefreshToken(cfg *config.Config, claims *jwtAuth.Claims) string {
	return jwtAuth.NewRefreshToken(cfg.Jwt.RefreshSecretKey, cfg.Jwt.RefreshDuration, &jwtAuth.Claims{
		UserID: claims.UserID,
	}).SignToken()
}

func (a *authRepository) GetOneUserByAuthProvider(authProviderName string, authProviderUserID string) (*entities.User, error) {
	user := &entities.User{}
	if err := a.db.Where(&entities.User{
		AuthProviderName:   authProviderName,
		AuthProviderUserID: authProviderUserID,
	},
	).First(user).Error; err != nil {
		log.Printf("error: GetOneUserByGoogleUserID: %s", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &entities.User{}, errors.New("error: user not found")
		}
		return &entities.User{}, errors.New("error: get user data failed")
	}
	return user, nil

}

func (a *authRepository) UpdateOneUser(in *entities.User) error {
	if err := a.db.Model(&in).Updates(in).Error; err != nil {
		log.Printf("error: UpdateOneUser: %s\n", err.Error())
		return errors.New("error: update user failed")
	}
	return nil
}
