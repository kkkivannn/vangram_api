package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	list, err := h.service.GetAllUsers(c)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}
