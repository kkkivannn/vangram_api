package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vangram_api/internal/http/middleware"
	"vangram_api/internal/service/post"
	"vangram_api/internal/service/user"
)

type request struct {
	PostID int `json:"post_id"`
}

type response struct {
	Post post.Post `json:"post"`
	User user.User `json:"user"`
}

func (h *Handler) getPost(c *gin.Context) {
	var r request
	err := c.BindJSON(&r)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	p, err := h.postService.GetPost(c, r.PostID)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	u, err := h.userService.GetUser(c, userID)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	p.User = u
	c.JSON(http.StatusOK, p)

}
