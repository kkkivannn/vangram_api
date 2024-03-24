package user

import (
	"context"
	"github.com/google/uuid"
	"time"
	t "vangram_api/pkg/tokens"
	"vangram_api/pkg/utils"
)

type StorageUser interface {
	CreateUser(ctx context.Context, user SaveUser) (int, error)
	ReadUser(ctx context.Context, id int) (User, error)
	ReadUserByNumber(ctx context.Context, number string) (int, error)
	UpdateUser(ctx context.Context, user SaveUser, userId int) error
	DeleteUser(ctx context.Context, id int) (string, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	SetUserRefreshToken(ctx context.Context, token string, userId int, sessionId string) error
	UpdateUserSession(ctx context.Context, sessionID string, userID int, oldRefreshToken string, newRefreshToken string) error
	DeleteUserSession(ctx context.Context, sessionId string, userID int) error
}

type ServiceUser struct {
	storage StorageUser
}

func New(storage StorageUser) *ServiceUser {
	return &ServiceUser{storage}
}

func (s *ServiceUser) CreateUser(ctx context.Context, user RequestUser) (int, error) {
	var saveUser SaveUser
	path, err := utils.SaveFile(user.Photo)
	if err != nil {
		return 0, err
	}
	saveUser = SaveUser{
		Name:       user.Name,
		Surname:    user.Surname,
		Age:        user.Age,
		Phone:      user.Phone,
		Photo:      path,
		CreatedAt:  user.CreatedAt,
		UploadedAt: user.UploadedAt,
	}
	return s.storage.CreateUser(ctx, saveUser)
}
func (s *ServiceUser) UpdateUser(ctx context.Context, user RequestUser, userId int) error {
	var saveUser SaveUser
	path, err := utils.SaveFile(user.Photo)
	if err != nil {
		return err
	}
	uploadedAt := time.Now()
	user.UploadedAt = &uploadedAt
	saveUser = SaveUser{
		Name:       user.Name,
		Surname:    user.Surname,
		Age:        user.Age,
		Phone:      user.Phone,
		Photo:      path,
		CreatedAt:  user.CreatedAt,
		UploadedAt: user.UploadedAt,
	}
	return s.storage.UpdateUser(ctx, saveUser, userId)
}

func (s *ServiceUser) DeleteUser(ctx context.Context, userId int) (string, error) {
	return s.storage.DeleteUser(ctx, userId)
}

func (s *ServiceUser) GetUser(ctx context.Context, userId int) (User, error) {
	return s.storage.ReadUser(ctx, userId)
}

func (s *ServiceUser) GetAllUsers(ctx context.Context) ([]User, error) {
	return s.storage.GetAllUsers(ctx)
}

func (s *ServiceUser) GenerateTokens(ctx context.Context, number string) (Tokens, error) {
	var tokens Tokens

	id, err := s.storage.ReadUserByNumber(ctx, number)

	if err != nil {
		return Tokens{}, err
	}
	sessionId := uuid.NewString()

	tokens.AccessToken, err = t.GenerateAccessToken(id, sessionId)
	if err != nil {
		return Tokens{}, err
	}

	tokens.RefreshToken, err = t.GenerateRefreshToken(sessionId, id)

	if err != nil {
		return Tokens{}, err
	}

	err = s.storage.SetUserRefreshToken(ctx, tokens.RefreshToken, id, sessionId)

	if err != nil {
		return Tokens{}, err
	}
	return tokens, nil
}

func (s *ServiceUser) RefreshTokens(ctx context.Context, requestRefreshToken string) (Tokens, error) {
	token, err := t.ParseRefreshToken(requestRefreshToken)
	if err != nil {
		return Tokens{}, err
	}
	accessToken, err := t.GenerateAccessToken(token.UserID, token.SessionID)
	if err != nil {
		return Tokens{}, err
	}
	refreshToken, err := t.GenerateRefreshToken(token.SessionID, token.UserID)
	if err != nil {
		return Tokens{}, err
	}

	err = s.storage.UpdateUserSession(ctx, token.SessionID, token.UserID, requestRefreshToken, refreshToken)
	if err != nil {
		return Tokens{}, err
	}
	return Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil

}

func (s *ServiceUser) RemoveUserSession(ctx context.Context, sessionId string, userID int) error {
	if err := s.storage.DeleteUserSession(ctx, sessionId, userID); err != nil {
		return err
	}
	return nil
}

func (s *ServiceUser) GetUserByNumber(ctx context.Context, number string) (User, error) {
	id, err := s.storage.ReadUserByNumber(ctx, number)
	if err != nil {
		return User{}, err
	}
	return s.storage.ReadUser(ctx, id)
}
