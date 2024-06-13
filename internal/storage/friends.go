package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"vangram_api/internal/service/friends"
	"vangram_api/internal/storage/postgres"
)

type Friends struct {
	db *pgxpool.Pool
}

func NewFriends(db *pgxpool.Pool) *Friends {
	return &Friends{db: db}
}

func (f *Friends) ReadAllFriends(ctx context.Context, userID int) ([]friends.Friend, error) {
	var allFriends []friends.Friend
	query := fmt.Sprintf("SELECT u.id, u.name, u.surname, u.age, u.photo FROM %s f INNER JOIN %s u ON f.friend_id = u.id WHERE f.user_id = $1", postgres.Friends, postgres.User)
	rows, err := f.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var friend friends.Friend
		if err := rows.Scan(&friend.ID, &friend.Name, &friend.Surname, &friend.Age, &friend.Photo); err != nil {
			return nil, err
		}
		if friend.Photo != nil {
			photo := fmt.Sprintf("%s%s", url, *friend.Photo)
			friend.Photo = &photo
		}
		allFriends = append(allFriends, friend)
	}
	return allFriends, nil
}

func (f *Friends) CreateFriend(ctx context.Context, userID, friendID int) error {
	query := fmt.Sprintf("INSERT INTO %s (friend_id, user_id) VALUES($1, $2)", postgres.Friends)
	_, err := f.db.Exec(ctx, query, friendID, userID)
	if err != nil {
		return err
	}
	return nil
}
