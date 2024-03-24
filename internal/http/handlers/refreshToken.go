package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) refreshToken(c *gin.Context) {
	var request refreshTokenRequest
	err := c.BindJSON(&request)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	tokens, err := h.userService.RefreshTokens(c, request.RefreshToken)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tokens)
	return
}
