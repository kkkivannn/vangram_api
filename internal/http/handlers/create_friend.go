package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vangram_api/internal/http/middleware"
)

type CreateFriendRequest struct {
	FriendID int `json:"friend_id"`
}

func (h *Handler) createFriend(ctx *gin.Context) {
	var createFriendRequest CreateFriendRequest
	err := ctx.BindJSON(&createFriendRequest)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	err = h.friendService.AddNewFriend(ctx, userID, createFriendRequest.FriendID)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.AbortWithStatus(http.StatusCreated)
	return

}
