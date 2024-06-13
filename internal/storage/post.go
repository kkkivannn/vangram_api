package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
	"vangram_api/internal/service/post"
	"vangram_api/internal/service/user"
	"vangram_api/internal/storage/postgres"
)

const url = "http://localhost:8080/image"

type PostStorage struct {
	db *pgxpool.Pool
}

func NewPostStorage(db *pgxpool.Pool) *PostStorage {
	return &PostStorage{db: db}
}

func (ps *PostStorage) CreatePost(ctx context.Context, post post.SavePost) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (photo, body, created_at, id_user) VALUES ($1, $2, $3, $4) returning id", postgres.Post)
	row := ps.db.QueryRow(ctx, query, post.Photo, post.Body, post.CreatedAt, post.UserID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (ps *PostStorage) ReadPost(ctx context.Context, postID int) (post.Post, error) {
	var p post.Post
	query := fmt.Sprintf("SELECT id, photo, count_likes, body, created_at, uploaded_at FROM %s WHERE id=$1", postgres.Post)
	err := ps.db.QueryRow(ctx, query, postID).Scan(&p.ID, &p.Photo, &p.CountLikes, &p.Body, &p.CreatedAt, &p.UploadedAt)
	if err != nil {
		return post.Post{}, err
	}
	p.Photo = fmt.Sprintf("%s%s", url, p.Photo)
	return p, nil
}

func (ps *PostStorage) ReadAllPosts(ctx context.Context) ([]post.Post, error) {
	var posts []post.Post
	query := fmt.Sprintf("SELECT p.id, p.photo, p.count_likes, p.body, p.created_at, p.uploaded_at, u.id, u.name, u.surname, u.photo FROM %s p INNER JOIN %s u on p.id_user=u.id ORDER BY p.created_at DESC", postgres.Post, postgres.User)
	rows, err := ps.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p post.Post
		var u user.User
		err := rows.Scan(&p.ID, &p.Photo, &p.CountLikes, &p.Body, &p.CreatedAt, &p.UploadedAt, &u.ID, &u.Name, &u.Surname, &u.Photo)
		if err != nil {
			return nil, err
		}
		p.Photo = fmt.Sprintf("%s%s", url, p.Photo)
		p.User = u
		if p.User.Photo != nil {
			userPhoto := fmt.Sprintf("%s%s", url, *p.User.Photo)
			p.User.Photo = &userPhoto
		}

		posts = append(posts, p)
	}
	return posts, nil

}

// TODO Добавить вместо метода триггер
func (ps *PostStorage) SetLike(ctx context.Context, postID int) error {
	query := fmt.Sprintf("UPDATE %s SET count_likes=count_likes + 1 WHERE id=$1", postgres.Post)
	rows, err := ps.db.Exec(ctx, query, postID)
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return err
	}
	return nil
}

func (ps *PostStorage) AddLikesPost(ctx context.Context, postID, userID, userPostID int, likedAt time.Time) error {
	query := fmt.Sprintf("INSERT INTO %s (id_user, id_post, liked_at, id_liked_user) VALUES ($1, $2, $3, $4)", postgres.LikeUsersPost)
	rows, err := ps.db.Exec(ctx, query, userID, postID, likedAt, userPostID)
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return err
	}
	return nil
}

func (ps *PostStorage) ReadLikesUserPosts(ctx context.Context, userID int) ([]post.Post, error) {
	var posts []post.Post
	query := fmt.Sprintf("SELECT p.id, p.photo, p.count_likes, p.body, p.created_at, p.uploaded_at, u.id, u.name, u.surname, u.photo, u.phone_number FROM %s lu INNER JOIN %s p on p.id=lu.id_post INNER JOIN %s u on u.id=lu.id_user WHERE lu.id_liked_user=$1", postgres.LikeUsersPost, postgres.Post, postgres.User)
	rows, err := ps.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p post.Post
		var u user.User
		err := rows.Scan(&p.ID, &p.Photo, &p.CountLikes, &p.Body, &p.CreatedAt, &p.UploadedAt,
			&u.ID, &u.Name, &u.Surname, &u.Photo, &u.Phone,
		)
		if err != nil {
			return nil, err
		}
		p.Photo = fmt.Sprintf("%s%s", url, p.Photo)
		p.User = u
		posts = append(posts, p)
	}
	return posts, nil
}

func (ps *PostStorage) ReadUserPosts(ctx context.Context, userID int) ([]post.Post, error) {
	var posts []post.Post
	query := fmt.Sprintf("SELECT p.id, p.photo, p.count_likes, p.body, p.created_at, p.uploaded_at, u.id, u.name, u.surname, u.photo, u.phone_number FROM %s p INNER JOIN %s u on u.id=p.id_user WHERE p.id_user=$1 ORDER BY p.created_at DESC", postgres.Post, postgres.User)
	rows, err := ps.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p post.Post
		var u user.User
		err := rows.Scan(&p.ID, &p.Photo, &p.CountLikes, &p.Body, &p.CreatedAt, &p.UploadedAt,
			&u.ID, &u.Name, &u.Surname, &u.Photo, &u.Phone,
		)
		if err != nil {
			return nil, err
		}
		p.Photo = fmt.Sprintf("%s%s", url, p.Photo)
		p.User = u
		if p.User.Photo != nil {
			userPhoto := fmt.Sprintf("%s%s", url, *p.User.Photo)
			p.User.Photo = &userPhoto
		}

		posts = append(posts, p)
	}
	return posts, nil
}
