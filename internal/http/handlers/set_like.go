package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
	"vangram_api/internal/http/middleware"
)

type SetLikeRequest struct {
	PostID     int `json:"post_id"`
	UserPostID int `json:"user_post_id"`
}

func (h *Handler) setLike(c *gin.Context) {
	var request SetLikeRequest
	err := c.BindJSON(&request)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	if err := h.postService.SetLikeToPost(c, request.PostID); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		slog.Error("Not Authorized")
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorized")
		return
	}
	if err := h.postService.AddLikesPost(c, request.PostID, userID, request.UserPostID, time.Now()); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
}
