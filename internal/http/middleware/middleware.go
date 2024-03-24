package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
	"vangram_api/pkg/tokens"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	userSession         = "sessionId"
)

func Identity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		slog.Error("Not Authorized")
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Вы не авторизованы")
		return
	}
	splitsHeader := strings.Split(header, " ")

	if len(splitsHeader) != 2 {
		slog.Error("Un valid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Не валидный токен")
		return
	}
	authorized, err := tokens.IsAuthorized(splitsHeader[1])
	if authorized {
		token, err := tokens.ParseAccessToken(splitsHeader[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"err": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set(userCtx, token.UserID)
		c.Set(userSession, token.SessionID)
		return
	}
	c.JSON(http.StatusUnauthorized, err.Error())
	c.Abort()
	return
}
func GetUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
func GetUserSessionID(c *gin.Context) (string, error) {
	id, ok := c.Get(userSession)
	if !ok {
		return "", errors.New("user session id not found")
	}

	idInt, ok := id.(string)
	if !ok {
		return "", errors.New("user session id is of invalid type")
	}

	return idInt, nil
}
