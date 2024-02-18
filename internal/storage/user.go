package storage

import (
	"context"
	"fmt"
	"strings"
	"vangram_api/internal/postgres"
	"vangram_api/internal/service/user"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(ctx context.Context, user user.RequestUser) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname) VALUES ($1, $2) returning id", postgres.Client)
	row := s.db.QueryRow(ctx, query, user.Name, user.Surname)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserStorage) ReadUser(ctx context.Context, id int) (user.User, error) {
	var resp user.User
	query := fmt.Sprintf("SELECT id, name, surname FROM %s WHERE id=$1", postgres.Client)
	err := s.db.QueryRow(ctx, query, id).Scan(&resp.ID, &resp.Name, &resp.Surname)
	if err != nil {
		return user.User{}, err
	}
	return resp, nil
}

func (s *UserStorage) UpdateUser(ctx context.Context, u user.User) ([]user.User, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	var newUsers []user.User
	if u.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *u.Name)
		argId++
	}
	if u.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *u.Surname)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", postgres.Client, setQuery, argId)
	args = append(args, u.ID)
	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return nil, err
		}
		return newUsers, err
	}

	queryUsers := fmt.Sprintf("SELECT id, name, surname FROM %s", postgres.Client)
	rows, err := s.db.Query(ctx, queryUsers)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	for rows.Next() {
		var u user.User
		err := rows.Scan(u.ID, u.Name, u.Surname)
		if err != nil {
			return nil, err
		}
		newUsers = append(newUsers, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return newUsers, nil
}

func (s *UserStorage) DeleteUser(ctx context.Context, id int) (string, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, postgres.Client)
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return "", err
	}
	return "User has been deleted", nil
}

func (s *UserStorage) GetAllUsers(ctx context.Context) ([]user.User, error) {
	var users []user.User
	query := fmt.Sprintf("SELECT id, name, surname FROM %s", postgres.Client)
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user user.User
		err := rows.Scan(&user.ID, &user.Name, &user.Surname)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
