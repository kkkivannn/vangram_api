package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vangram_api/internal/lib/api/response"
)

type GetAllUsers interface {
	GetAllUsers(ctx context.Context) ([]response.UserResponse, error)
}

func (h *Handlers) getAllUsers(c *gin.Context) {
	list, err := h.services.GetAllUsers(c)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}
