package server

import (
	"github.com/kritAsawaniramol/myAnimeList-server/module/animeList/animeListHandler"
	"github.com/kritAsawaniramol/myAnimeList-server/module/animeList/animeListRepository"
	"github.com/kritAsawaniramol/myAnimeList-server/module/animeList/animeListUsecase"
)

func (s *ginServer) animeListService() {
	repo := animeListRepository.NewAnimeRepository(s.db)
	usecase := animeListUsecase.NewAnimeListUsecase(repo)
	handler := animeListHandler.NewAnimeListHandler(usecase)

	s.app.POST("/animeList", s.middleware.JwtAuthorization(), handler.AddAnimeToMyList)
	s.app.GET("/animeList", s.middleware.JwtAuthorization(), handler.GetAnimList)
	s.app.GET("/animeList/:malID", s.middleware.JwtAuthorization(), handler.GetOneAnimeList)
	s.app.PATCH("/animeList/:malID", s.middleware.JwtAuthorization(), handler.UpdateOneAnimeList)
	s.app.DELETE("/animeList/:malID", s.middleware.JwtAuthorization(), handler.RemoveAnimeInAnimeList)
}
