package commentHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritAsawaniramol/myAnimeList-server/module/comment"
	"github.com/kritAsawaniramol/myAnimeList-server/module/comment/commentUsecase"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/request"
)

type (
	CommentHttpHandler interface {
		PostComment(ctx *gin.Context)
		GetAnimeCommentsReq(ctx *gin.Context)
	}

	commentHttpHandler struct {
		commentUsecase commentUsecase.CommentUsecase
	}
)

// GetCommentsByMalID implements CommentHttpHandler.
func (h *commentHttpHandler) GetAnimeCommentsReq(ctx *gin.Context) {
	malID := ctx.Query("mal_id")
	if malID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error: require mal_id in query"})
		return
	}

	comments, err := h.commentUsecase.GetCommentsByMalID(malID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func NewCommentHttpHandler(commentUsecase commentUsecase.CommentUsecase) CommentHttpHandler {
	return &commentHttpHandler{commentUsecase: commentUsecase}
}

// PostComment implements CommentHttpHandler.
func (h *commentHttpHandler) PostComment(ctx *gin.Context) {
	wrapper := request.ContextWrapper(ctx)
	commentReq := &comment.PostCommentReq{}
	if err := wrapper.Bind(commentReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	commentReq.UserID = ctx.GetUint("userID")
	if err := h.commentUsecase.PostComment(commentReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "post comment successed"})
}
