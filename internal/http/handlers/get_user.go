package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestGetUser struct {
	ID int `json:"id"`
}

func (h *Handler) getUser(c *gin.Context) {
	var request RequestGetUser
	err := c.BindJSON(&request)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.userService.GetUser(c, request.ID)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
