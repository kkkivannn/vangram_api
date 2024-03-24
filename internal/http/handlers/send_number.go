package handlers

import (
	"log/slog"
	"net/http"
	"vangram_api/internal/service/user"

	"github.com/gin-gonic/gin"
)

func (h *Handler) sendNumber(c *gin.Context) {
	var requestUser user.SendNumberRequest

	if err := c.BindJSON(&requestUser); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	//if err := c.ShouldBindWith(&requestUser, binding.FormMultipart); err != nil {
	//	if requestUser.Photo == nil {
	//		fmt.Println("true")
	//	}
	//	slog.Error(err.Error())
	//	c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//requestUser.CreatedAt = time.Now()
	//_, err := h.userService.CreateUser(c, requestUser)
	//if err != nil {
	//	slog.Error(err.Error())
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//tokens, err := h.userService.GenerateTokens(c, requestUser.Phone)
	//if err != nil {
	//	slog.Error(err.Error())
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	return
}
