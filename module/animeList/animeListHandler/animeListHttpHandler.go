package animeListHandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritAsawaniramol/myAnimeList-server/module/animeList"
	"github.com/kritAsawaniramol/myAnimeList-server/module/animeList/animeListUsecase"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/request"
	"github.com/kritAsawaniramol/myAnimeList-server/util"
)

type (
	AnimeListHandler interface {
		AddAnimeToMyList(ctx *gin.Context)
		GetOneAnimeList(ctx *gin.Context)
		UpdateOneAnimeList(ctx *gin.Context)
		GetAnimList(ctx *gin.Context)
		RemoveAnimeInAnimeList(ctx *gin.Context)
	}

	animeHttpHandler struct {
		animeListUsecase animeListUsecase.AnimeListUsecase
	}
)

// RemoveAnimeInAnimeList implements AnimeListHandler.
func (a *animeHttpHandler) RemoveAnimeInAnimeList(ctx *gin.Context) {
	malID := ctx.Param("malID")
	if err := a.animeListUsecase.RemoveOneAnimeInList(malID, ctx.GetUint("userID")); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// GetAnimList implements AnimeListHandler.
func (a *animeHttpHandler) GetAnimList(ctx *gin.Context) {
	res, err := a.animeListUsecase.GetAnimeList(ctx.GetUint("userID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// UpdateOneAnimeList implements AnimeListHandler.
func (a *animeHttpHandler) UpdateOneAnimeList(ctx *gin.Context) {
	malID := ctx.Param("malID")
	wrapper := request.ContextWrapper(ctx)
	req := &animeList.UpdateOneAnimeListReq{}
	if err := wrapper.Bind(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	util.PrintObjInJson(req)
	req.MalID = malID
	req.UserID = ctx.GetUint("userID")
	res, err := a.animeListUsecase.UpdateOneAnimeList(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	util.PrintObjInJson(res)
	ctx.JSON(http.StatusOK, gin.H{"anime_list": res})
}

// GetOneAnimeList implements AnimeListHandler.
func (a *animeHttpHandler) GetOneAnimeList(ctx *gin.Context) {
	malID := ctx.Param("malID")
	res, err := a.animeListUsecase.GetOneAnimeList(malID, ctx.GetUint("userID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// AddAnimeToMyList implements AnimeListHandler.
func (a *animeHttpHandler) AddAnimeToMyList(ctx *gin.Context) {
	wrapper := request.ContextWrapper(ctx)
	addAnimeToMyListReq := &animeList.AddAnimeToMyListReq{}
	if err := wrapper.Bind(addAnimeToMyListReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	addAnimeToMyListReq.UserID = ctx.GetUint("userID")
	fmt.Printf("addAnimeToMyListReq: %v\n", addAnimeToMyListReq)
	err := a.animeListUsecase.AddAnimeToMyList(addAnimeToMyListReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func NewAnimeListHandler(animeListUsecase animeListUsecase.AnimeListUsecase) AnimeListHandler {
	return &animeHttpHandler{
		animeListUsecase: animeListUsecase,
	}
}
