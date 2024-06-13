package post

import (
	"context"
	"time"
	"vangram_api/pkg/utils"
)

type StoragePost interface {
	CreatePost(ctx context.Context, post SavePost) (int, error)
	ReadPost(ctx context.Context, postID int) (Post, error)
	ReadAllPosts(ctx context.Context) ([]Post, error)
	ReadUserPosts(ctx context.Context, userID int) ([]Post, error)
	SetLike(ctx context.Context, postID int) error
	AddLikesPost(ctx context.Context, postID, userID, userPostID int, likedAt time.Time) error
	ReadLikesUserPosts(ctx context.Context, userID int) ([]Post, error)
}

type Service struct {
	storagePost StoragePost
}

func NewService(storage StoragePost) *Service {
	return &Service{storagePost: storage}

}

func (s *Service) CreateUserPost(ctx context.Context, post CreatePostModel) (int, error) {

	var savePost SavePost

	path, err := utils.SaveFile(post.Photo)

	if err != nil {
		return 0, err
	}

	savePost = SavePost{
		Photo:     path,
		Body:      post.Body,
		CreatedAt: time.Now(),
		UserID:    post.UserID,
	}

	return s.storagePost.CreatePost(ctx, savePost)
}

func (s *Service) GetPost(ctx context.Context, postID int) (Post, error) {
	return s.storagePost.ReadPost(ctx, postID)
}

func (s *Service) GetAllPosts(ctx context.Context) ([]Post, error) {
	return s.storagePost.ReadAllPosts(ctx)
}

func (s *Service) SetLikeToPost(ctx context.Context, postID int) error {
	return s.storagePost.SetLike(ctx, postID)
}

func (s *Service) AddLikesPost(ctx context.Context, postID, userID, userPostID int, likedAt time.Time) error {
	return s.storagePost.AddLikesPost(ctx, postID, userID, userPostID, likedAt)
}

func (s *Service) GetLikesUsersPosts(ctx context.Context, userID int) ([]Post, error) {
	return s.storagePost.ReadLikesUserPosts(ctx, userID)
}

func (s *Service) GetUserPosts(ctx context.Context, userID int) ([]Post, error) {
	return s.storagePost.ReadUserPosts(ctx, userID)
}
