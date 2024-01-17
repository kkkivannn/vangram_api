package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"vangram_api/utils"
)

func (h *MainHandlers) getUser(c *gin.Context) {
	var inputUser utils.UserDto
	err := c.BindJSON(&inputUser)
	if err != nil {
		logrus.Fatal(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.GetUser(inputUser.Id)
	if err != nil {
		logrus.Fatal(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *MainHandlers) deleteUser(c *gin.Context) {
	var inputUser utils.UserDto
	if err := c.BindJSON(&inputUser); err != nil {
		logrus.Fatal(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.services.RemoveUser(inputUser.Id)
	if err != nil {
		logrus.Fatal(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
