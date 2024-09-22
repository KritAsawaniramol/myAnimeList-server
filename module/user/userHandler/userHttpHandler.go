package userHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritAsawaniramol/myAnimeList-server/module/user/userUsecase"
)

type (
	UserHttpHandler interface {
		GetUserProfile(ctx *gin.Context)
	}

	userHttpHandler struct {
		userUsecase userUsecase.UserUsecase
	}
)

// GetUserProfile implements UserHttpHandler.
func (u *userHttpHandler) GetUserProfile(ctx *gin.Context) {
	user, err := u.userUsecase.GetUserProfile(ctx.GetUint("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func NewUserHttpHandler(userUsecase userUsecase.UserUsecase) UserHttpHandler {
	return &userHttpHandler{userUsecase: userUsecase}
}
