package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
	"vangram_api/internal/database"
	"vangram_api/internal/handlers"
	"vangram_api/internal/lib/api/response"
)

type AuthorizeRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthorizeRepository {
	return &AuthorizeRepository{db: db}
}

func (ar *AuthorizeRepository) Create(ctx context.Context, user handlers.RequestCreateUser) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name, surname) VALUES ($1, $2) returning id", database.Client)
	row := ar.db.QueryRow(ctx, query, user.Name, user.Surname)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (ar *AuthorizeRepository) Read(ctx context.Context, id int) (response.UserResponse, error) {
	var user response.UserResponse
	query := fmt.Sprintf("select id, name, surname from %s where id=$1", database.Client)
	err := ar.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Surname)
	if err != nil {
		return response.UserResponse{}, err
	}
	return user, nil
}

func (ar *AuthorizeRepository) Update(ctx context.Context, user handlers.RequestUpdateUser) ([]handlers.RequestUpdateUser, error) {
	tx, err := ar.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	var newUsers []handlers.RequestUpdateUser
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
	_, err = ar.db.Exec(ctx, query, args...)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return nil, err
		}
		return newUsers, err
	}

	queryUsers := fmt.Sprintf("SELECT id, name, surname FROM %s", database.Client)
	rows, err := ar.db.Query(ctx, queryUsers)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	for rows.Next() {
		var user handlers.RequestUpdateUser
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

func (ar *AuthorizeRepository) Delete(ctx context.Context, id int) (string, error) {
	query := fmt.Sprintf(`delete from %s where id=$1`, database.Client)
	_, err := ar.db.Exec(ctx, query, id)
	if err != nil {
		return "", err
	}
	return "User has been deleted", nil
}

func (ar *AuthorizeRepository) GetAll(ctx context.Context) ([]response.UserResponse, error) {
	var users []response.UserResponse
	query := fmt.Sprintf("SELECT id, name, surname FROM %s", database.Client)
	rows, err := ar.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user response.UserResponse
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
