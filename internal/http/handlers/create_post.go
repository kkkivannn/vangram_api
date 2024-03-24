package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log/slog"
	"net/http"
	"vangram_api/internal/http/middleware"
	"vangram_api/internal/service/post"
)

func (h *Handler) createPost(c *gin.Context) {
	var request post.CreatePostModel
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	userID, err := middleware.GetUserID(c)

	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	request.UserID = userID

	id, err := h.postService.CreateUserPost(c, request)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
