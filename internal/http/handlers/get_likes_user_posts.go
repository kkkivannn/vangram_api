package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vangram_api/internal/http/middleware"
)

func (h *Handler) getLikesUserPosts(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		slog.Error("Not Authorized")
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Not Authorized",
		})
		return
	}
	posts, err := h.postService.GetLikesUsersPosts(c, userID)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}
