package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestDeleteUser struct {
	ID int `json:"id"`
}

func (h *Handler) deleteUser(c *gin.Context) {
	var request RequestDeleteUser
	if err := c.BindJSON(&request); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	status, err := h.userService.DeleteUser(c, request.ID)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
	return
}
