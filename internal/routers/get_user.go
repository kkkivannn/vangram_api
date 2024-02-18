package routers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type RequestGetUser struct {
	ID int `json:"id"`
}

func (r *Route) getUser(c *gin.Context) {
	var request RequestGetUser
	err := c.BindJSON(&request)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := r.service.GetUser(c, request.ID)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
