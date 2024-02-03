package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"vangram_api/utils"
)

func (h *MainHandlers) signUp(c *gin.Context) {
	inputUser := &utils.UserDTO{}
	if err := c.BindJSON(inputUser); err != nil {
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

func (h *MainHandlers) signIn(c *gin.Context) {
	inputUser := &utils.InputUser{}
	if err := c.BindJSON(inputUser); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

}
