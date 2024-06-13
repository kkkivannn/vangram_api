package tokens

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log/slog"
	"time"
)

const (
	Salt                   = "dslkj932q90jdqos0219jd3fjreasokcmnurn4875678f"
	AccessTokenSigningKey  = "WnRcez^SQQi_['q*3SJmR|5^P4-~ga"
	RefreshTokenSigningKey = "LloVYPiV5km9FP}9>;Fvf}%(Z}B|]="
	AccessTokenTTL         = 50 * time.Minute
	RefreshTokenTTL        = 720 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type Tokens struct {
	UserID    int
	SessionID string
}

func ParseAccessToken(requestToken string) (Tokens, error) {
	token, err := jwt.ParseWithClaims(requestToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Error("Not valid token")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AccessTokenSigningKey), nil
	})
	if err != nil {
		return Tokens{}, err
	}
	if !token.Valid {
		return Tokens{}, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return Tokens{}, err
	}
	return Tokens{
		UserID:    claims.UserId,
		SessionID: claims.Id,
	}, nil
}
func ParseRefreshToken(requestToken string) (Tokens, error) {
	token, err := jwt.ParseWithClaims(requestToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Error("Not valid token")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(RefreshTokenSigningKey), nil
	})
	if err != nil {
		return Tokens{}, err
	}
	if !token.Valid {
		return Tokens{}, errors.New("Token not valid")
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return Tokens{}, errors.New("Token claims are not of type *refreshTokenClaims")
	}
	tokenData := Tokens{
		SessionID: claims.Id,
		UserID:    claims.UserId,
	}
	return tokenData, nil
}

func IsAuthorized(requestToken string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Error("Not valid token")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AccessTokenSigningKey), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateAccessToken(userId int, sessionID string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        sessionID,
		},
		UserId: userId,
	})
	return accessToken.SignedString([]byte(AccessTokenSigningKey))
}

func GenerateRefreshToken(sessionId string, userId int) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        sessionId,
		},
		UserId: userId,
	})
	return refreshToken.SignedString([]byte(RefreshTokenSigningKey))
}
