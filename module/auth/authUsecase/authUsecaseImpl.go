package authUsecase

import (
	"log"

	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth/authRepository"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/jwtAuth"
)

type (
	authUsecaseImpl struct {
		authRepository authRepository.AuthServiceRepository
	}
)

// RefreshToken implements AuthUsecase.
func (a *authUsecaseImpl) RefreshToken(cfg *config.Config, credentialID uint, refreshToken string) (*auth.CredentialRes, error) {

	//PareseToken
	claims, err := jwtAuth.ParseToken(cfg.Jwt.RefreshSecretKey, refreshToken)
	if err != nil {
		log.Printf("error: RefreshToken: %s\n", err.Error())
		return nil, err
	}

	user, err := a.authRepository.GetOneUserByID(claims.UserID)
	if err != nil {
		log.Printf("error: RefreshToken: %s\n", err.Error())
		return nil, err
	}

	accessToken := jwtAuth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtAuth.Claims{
		UserID: user.ID,
	}).SignToken()

	refreshToken = jwtAuth.ReloadToken(cfg.Jwt.RefreshSecretKey, claims.ExpiresAt.Unix(), &jwtAuth.Claims{
		UserID: user.ID,
	})

	updateCondition := &entities.Credential{}
	updateCondition.ID = credentialID
	updateCondition.UserID = claims.UserID
	updateCondition.RefreshToken = refreshToken
	if err := a.authRepository.UpdateOneUserCredential(updateCondition, &entities.Credential{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}); err != nil {
		return nil, err
	}

	credential, err := a.authRepository.GetOneUserCredentialByID(credentialID)
	if err != nil {
		return nil, err
	}

	return &auth.CredentialRes{
		Id:           credential.ID,
		UserID:       credential.UserID,
		AccessToken:  credential.AccessToken,
		RefreshToken: credential.RefreshToken,
		CreatedAt:    credential.CreatedAt,
		UpdatedAt:    credential.UpdatedAt,
	}, nil
}

// Logout implements AuthUsecase.
func (a *authUsecaseImpl) Logout(credentialID uint) error {
	if err := a.authRepository.DeleteOneUserCredentialByID(credentialID); err != nil {
		return err
	}
	return nil
}

// LoginWithOauth implements AuthUsecase.
func (a *authUsecaseImpl) LoginWithOauth(cfg *config.Config, in *auth.LoginWithOauth) (*auth.LoginWithOauthRes, error) {
	user, err := a.authRepository.GetOneUserByAuthProvider(in.AuthProviderName, in.AuthProviderUserID)
	if err != nil {
		if err.Error() == "error: user not found" {
			_, createUserErr := a.authRepository.CreateOneUser(&entities.User{
				Name:               in.Name,
				Email:              in.Email,
				AvatarURL:          in.AvatarURL,
				AuthProviderName:   in.AuthProviderName,
				AuthProviderUserID: in.AuthProviderUserID,
			})
			if createUserErr != nil {
				return &auth.LoginWithOauthRes{}, err
			}
			newUser, err := a.authRepository.GetOneUserByAuthProvider(in.AuthProviderName, in.AuthProviderUserID)
			if err != nil {
				return &auth.LoginWithOauthRes{}, err
			}
			user = newUser
		} else {
			return &auth.LoginWithOauthRes{}, err
		}
	} else {
		if user.Name != in.Name || user.AvatarURL != in.AvatarURL {
			user.Name = in.Name
			user.AvatarURL = in.AvatarURL
			_ = a.authRepository.UpdateOneUser(user)
		}
	}

	accessToken := a.authRepository.AccessToken(cfg, &jwtAuth.Claims{
		UserID: user.ID,
	})

	refreshToken := a.authRepository.RefreshToken(cfg, &jwtAuth.Claims{
		UserID: user.ID,
	})

	credentialID, err := a.authRepository.InserOneUserCredential(&entities.Credential{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		return &auth.LoginWithOauthRes{}, err
	}

	credential, err := a.authRepository.GetOneUserCredentialByID(credentialID)
	if err != nil {
		return &auth.LoginWithOauthRes{}, err
	}

	return &auth.LoginWithOauthRes{
		LoginWithOauth: auth.LoginWithOauth{
			Email:              user.Email,
			Name:               user.Name,
			AuthProviderName:   user.AuthProviderName,
			AuthProviderUserID: user.AuthProviderUserID,
			AvatarURL:          in.AvatarURL,
		},
		Credential: &auth.CredentialRes{
			Id:           credential.ID,
			UserID:       credential.UserID,
			AccessToken:  credential.AccessToken,
			RefreshToken: credential.RefreshToken,
			CreatedAt:    credential.CreatedAt,
			UpdatedAt:    credential.UpdatedAt,
		},
	}, nil
}

func NewAuthUsecaseImpl(authRepository authRepository.AuthServiceRepository) AuthUsecase {
	return &authUsecaseImpl{
		authRepository: authRepository,
	}
}
