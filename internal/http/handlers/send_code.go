package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
	"vangram_api/internal/service/user"
)

func (h *Handler) sendCode(c *gin.Context) {
	var input user.SignInUserRequest

	if err := c.BindJSON(&input); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if input.Code != "1111" {
		err := errors.New("Код не верный")
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.userService.GetUserByNumber(c, input.Phone)
	if err != nil {
		var saveUser user.RequestUser
		saveUser.Phone = input.Phone
		saveUser.CreatedAt = time.Now()
		_, err := h.userService.CreateUser(c, saveUser)
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		tokens, err := h.userService.GenerateTokens(c, input.Phone)
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusCreated, map[string]interface{}{
			"has_profile":   false,
			"access_token":  tokens.AccessToken,
			"refresh_token": tokens.RefreshToken,
		})
		return
	}
	tokens, err := h.userService.GenerateTokens(c, input.Phone)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"has_profile":   true,
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
	return
}
