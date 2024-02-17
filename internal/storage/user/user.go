package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
	"vangram_api/internal/database"
	"vangram_api/internal/service"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateUser(ctx context.Context, user service.RequestUser) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname) VALUES ($1, $2) returning id", database.Client)
	row := s.db.QueryRow(ctx, query, user.Name, user.Surname)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) ReadUser(ctx context.Context, id int) (service.User, error) {
	var user service.User
	query := fmt.Sprintf("SELECT id, name, surname FROM %s WHERE id=$1", database.Client)
	err := s.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Surname)
	if err != nil {
		return service.User{}, err
	}
	return user, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user service.User) ([]service.User, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	var newUsers []service.User
	if user.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *user.Name)
		argId++
	}
	if user.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *user.Surname)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", database.Client, setQuery, argId)
	args = append(args, user.ID)
	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return nil, err
		}
		return newUsers, err
	}

	queryUsers := fmt.Sprintf("SELECT id, name, surname FROM %s", database.Client)
	rows, err := s.db.Query(ctx, queryUsers)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	for rows.Next() {
		var user service.User
		err := rows.Scan(user.ID, user.Name, user.Surname)
		if err != nil {
			return nil, err
		}
		newUsers = append(newUsers, user)
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

func (s *Storage) DeleteUser(ctx context.Context, id int) (string, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, database.Client)
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return "", err
	}
	return "User has been deleted", nil
}

func (s *Storage) GetAllUsers(ctx context.Context) ([]service.User, error) {
	var users []service.User
	query := fmt.Sprintf("SELECT id, name, surname FROM %s", database.Client)
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user service.User
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
