package server

import (
	"github.com/kritAsawaniramol/myAnimeList-server/module/comment/commentHandler"
	"github.com/kritAsawaniramol/myAnimeList-server/module/comment/commentRepository"
	"github.com/kritAsawaniramol/myAnimeList-server/module/comment/commentUsecase"
)

func (s *ginServer) commentService() {
	repo := commentRepository.NewPostgresRepository(s.db)
	usecase := commentUsecase.NewCommentUsecase(repo)
	handler := commentHandler.NewCommentHttpHandler(usecase)

	s.app.POST("/comment", s.middleware.JwtAuthorization(), handler.PostComment)
	s.app.GET("/comment", handler.GetAnimeCommentsReq)
}
