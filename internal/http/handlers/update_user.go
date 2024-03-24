package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log/slog"
	"net/http"
	"vangram_api/internal/http/middleware"
	"vangram_api/internal/service/user"
)

func (h *Handler) updateUser(c *gin.Context) {
	var input user.RequestUser
	if err := c.ShouldBindWith(&input, binding.FormMultipart); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userId, err := middleware.GetUserID(c)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := h.userService.UpdateUser(c, input, userId); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//u, err := h.userService.GetUser(c, userId)
	//tokens, err := h.userService.GenerateTokens(c, u.Phone)
	//if err != nil {
	//	slog.Error(err.Error())
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//c.JSON(http.StatusOK, tokens)
	return
}
