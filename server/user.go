package server

import (
	"github.com/kritAsawaniramol/myAnimeList-server/module/user/userHandler"
	"github.com/kritAsawaniramol/myAnimeList-server/module/user/userRepository"
	"github.com/kritAsawaniramol/myAnimeList-server/module/user/userUsecase"
)

func (s *ginServer) userService() {
	repo := userRepository.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repo)
	handler := userHandler.NewUserHttpHandler(usecase)

	s.app.GET("/profile", s.middleware.JwtAuthorization(), handler.GetUserProfile)
}
