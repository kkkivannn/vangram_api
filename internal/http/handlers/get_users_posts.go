package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vangram_api/internal/http/middleware"
)

func (h *Handler) getUsersPosts(ctx *gin.Context) {
	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		slog.Error("get user ID failed", err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	posts, err := h.postService.GetUserPosts(ctx, userID)
	if err != nil {
		slog.Error("get user posts failed", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}
