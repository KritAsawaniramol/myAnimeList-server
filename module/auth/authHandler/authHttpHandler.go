package authHandler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth/authUsecase"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/request"
	"github.com/markbates/goth/gothic"
)

type (
	AuthHttpHandler interface {
		GetOauth(ctx *gin.Context)
		GetOauthCallback(ctx *gin.Context)
		Logout(ctx *gin.Context)
		RefreshToken(ctx *gin.Context)
	}

	authHttpHandler struct {
		authUsecase authUsecase.AuthUsecase
		cfg         *config.Config
	}
)

// RefreshToken implements AuthHttpHandler.
func (a *authHttpHandler) RefreshToken(ctx *gin.Context) {
	// credential_id and refreshToken
	wrapper := request.ContextWrapper(ctx)
	refreshTokenReq := &auth.RefreshTokenReq{}
	if err := wrapper.Bind(refreshTokenReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := a.authUsecase.RefreshToken(a.cfg, refreshTokenReq.CredentialID, refreshTokenReq.RefreshToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func NewAuthHttpHandler(authUsecase authUsecase.AuthUsecase, cfg *config.Config) AuthHttpHandler {
	return &authHttpHandler{
		authUsecase: authUsecase,
		cfg:         cfg,
	}
}

// GetOauth implements AuthHttpHandler.
func (a *authHttpHandler) GetOauth(ctx *gin.Context) {
	provider := ctx.Param("provider")
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "provider", provider))
	_, err := gothic.GetAuthURL(ctx.Writer, ctx.Request)
	if err != nil {
		log.Printf("err: %s\n", err.Error())
	}

	if _, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request); err == nil {
		http.Redirect(ctx.Writer, ctx.Request, fmt.Sprintf("http://%s:%d/auth/%s/callback", a.cfg.App.Host, a.cfg.App.Port, provider), http.StatusFound)
	} else {
		gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
	}
}

// GetOauthCallback implements AuthHttpHandler.
func (a *authHttpHandler) GetOauthCallback(ctx *gin.Context) {
	provider := ctx.Param("provider")

	if ctx.Query("error") != "" {
		ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("http://%s:%d", a.cfg.Client.Host, a.cfg.Client.Port))
	}

	ctx.Request = ctx.Request.WithContext(
		context.WithValue(ctx.Request.Context(), "provider", provider))

	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintln(ctx.Writer, err)
		return
	}

	loginWithOauth := &auth.LoginWithOauth{
		Email:              user.Email,
		Name:               user.Name,
		AuthProviderName:   provider,
		AuthProviderUserID: user.UserID,
		AvatarURL:          user.AvatarURL,
	}

	res, err := a.authUsecase.LoginWithOauth(a.cfg, loginWithOauth)
	fmt.Printf("res: %v\n", res)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("http://%s:%d", a.cfg.Client.Host, a.cfg.Client.Port))
		return
	}

	ctx.SetCookie("access_token", res.Credential.AccessToken, 90, "/", a.cfg.Client.Host, false, false)

	ctx.SetCookie("refresh_token", res.Credential.RefreshToken, 90, "/", a.cfg.Client.Host, false, false)

	ctx.SetCookie("credential_id", fmt.Sprintf("%d", res.Credential.Id), 90, "/", a.cfg.Client.Host, false, false)

	ctx.SetCookie("is_user_logged_in", "1", int(a.cfg.Jwt.RefreshDuration), "/", a.cfg.Client.Host, false, false)

	ctx.Redirect(http.StatusFound, fmt.Sprintf("http://%s:%d", a.cfg.Client.Host, a.cfg.Client.Port))

}

// Logout implements AuthHttpHandler.
func (a *authHttpHandler) Logout(ctx *gin.Context) {
	StrCredentialID, err := ctx.Cookie("credential_id")
	if err != nil {
		log.Printf("error: Logout: %s\n", err.Error())
	}
	U64Credential, err := strconv.ParseUint(StrCredentialID, 10, 64)
	if err != nil {
		log.Printf("error: Logout: %s\n", err.Error())
	}
	if err := a.authUsecase.Logout(uint(U64Credential)); err != nil {
		log.Printf("error: Logout: %s\n", err.Error())
	}

	ctx.SetCookie("access_token", "", -1, "/", a.cfg.Client.Host, false, false)
	ctx.SetCookie("credential_id", "", -1, "/", a.cfg.Client.Host, false, false)
	ctx.SetCookie("refresh_token", "", -1, "/", a.cfg.Client.Host, false, false)
	ctx.SetCookie("is_user_logged_in", "", -1, "/", a.cfg.Client.Host, false, false)

	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
