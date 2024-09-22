package server

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth/authHandler"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth/authRepository"
	"github.com/kritAsawaniramol/myAnimeList-server/module/auth/authUsecase"
	"github.com/kritAsawaniramol/myAnimeList-server/util"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func (s *ginServer) authService() {
	store := sessions.NewCookieStore([]byte(s.cfg.Sessions.Secret))
	store.MaxAge(s.cfg.Sessions.MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = func() bool { return (s.cfg.App.Stage == "prod") }()
	gothic.Store = store

	util.PrintObjInJson(s.cfg.Facebook)

	goth.UseProviders(
		google.New(
			s.cfg.Google.ClientID,
			s.cfg.Google.ClientSecret,
			fmt.Sprintf("http://%s:%d/auth/google/callback", s.cfg.App.Host, s.cfg.App.Port),
			"email",
			"profile",
		),
		facebook.New(
			s.cfg.Facebook.ClientID,
			s.cfg.Facebook.ClientSecret,
			fmt.Sprintf("http://%s:%d/auth/facebook/callback", s.cfg.App.Host, s.cfg.App.Port),
			
		),
	)

	authRepo := authRepository.NewAuthRepository(s.db)

	authUsecase := authUsecase.NewAuthUsecaseImpl(authRepo)

	handler := authHandler.NewAuthHttpHandler(authUsecase, s.cfg)

	s.app.GET("/auth/:provider", handler.GetOauth)
	s.app.GET("/auth/:provider/callback", handler.GetOauthCallback)
	s.app.POST("/auth/refresh-token", handler.RefreshToken)
	s.app.POST("/auth/logout", handler.Logout)
}
