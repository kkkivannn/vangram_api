package handlers

import (
	"log/slog"
	"net/http"
	"vangram_api/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}
	list, err := h.userService.GetAllUsers(c, userID)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
	return
}
