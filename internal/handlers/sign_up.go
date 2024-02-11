package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CreateUser interface {
	CreateUser(ctx context.Context, user RequestCreateUser) (int, error)
}

type RequestCreateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (h *Handlers) signUp(c *gin.Context) {
	var inputUser RequestCreateUser
	if err := c.BindJSON(&inputUser); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(c, inputUser)
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
