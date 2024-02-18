package handlers

import (
	"net/http"
	"vangram_api/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RequestCreateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (r *Route) signUp(c *gin.Context) {
	var inputUser user.RequestUser
	if err := c.BindJSON(&inputUser); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := r.service.CreateUser(c, inputUser)
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
