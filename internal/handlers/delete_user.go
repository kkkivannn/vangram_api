package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type DeleteUser interface {
	DeleteUser(ctx context.Context, id int) (string, error)
}

type RequestDeleteUser struct {
	ID int `json:"id"`
}

func (h *Handlers) deleteUser(c *gin.Context) {
	var request RequestDeleteUser
	if err := c.BindJSON(&request); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.services.DeleteUser(c, request.ID)
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
