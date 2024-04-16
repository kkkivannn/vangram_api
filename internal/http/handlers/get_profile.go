package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vangram_api/internal/http/middleware"
)

func (h *Handler) getProfile(c *gin.Context) {
	id, err := middleware.GetUserID(c)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}
	user, err := h.userService.GetUser(c, id)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
