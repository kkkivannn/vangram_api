package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vangram_api/internal/lib/api/response"
)

type GetUser interface {
	GetUser(ctx context.Context, id int) (response.UserResponse, error)
}

type RequestGetUser struct {
	ID int `json:"id"`
}

func (h *Handlers) getUser(c *gin.Context) {
	var request RequestGetUser
	err := c.BindJSON(&request)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.GetUser(c, request.ID)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
