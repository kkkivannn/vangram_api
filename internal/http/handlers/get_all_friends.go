package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vangram_api/internal/http/middleware"
)

func (h *Handler) getAllFriends(ctx *gin.Context) {
	userID, err := middleware.GetUserID(ctx)

	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	friends, err := h.friendService.GetAllFriends(ctx, userID)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, friends)
}
