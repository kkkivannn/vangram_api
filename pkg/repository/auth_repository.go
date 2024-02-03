package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
	"vangram_api/pkg/database"
	"vangram_api/utils"
)

type AuthRepositoryProvider interface {
	Create(ctx context.Context, user *utils.UserDTO) (int, error)
	Read(ctx context.Context, id int) (utils.UserDTO, error)
	Update(ctx context.Context, user *utils.UserDTO) ([]utils.UserDTO, error)
	Delete(ctx context.Context, id int) (string, error)
}

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuth(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) Create(ctx context.Context, user *utils.UserDTO) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name, surname) VALUES ($1, $2) returning id", database.Client)
	row := ar.db.QueryRow(ctx, query, user.Name, user.Surname)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (ar *AuthRepository) Read(ctx context.Context, id int) (utils.UserDTO, error) {
	var user utils.UserDTO
	query := fmt.Sprintf("select id, name, surname from %s where id=$1", database.Client)
	err := ar.db.QueryRow(ctx, query, id).Scan(&user.Id, &user.Name, &user.Surname)
	if err != nil {
		return utils.UserDTO{}, err
	}
	return user, nil
}

func (ar *AuthRepository) Update(ctx context.Context, user *utils.UserDTO) ([]utils.UserDTO, error) {
	tx, err := ar.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	var newUsers []utils.UserDTO
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
	//setValues == "name=1, surname=2"

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", database.Client, setQuery, argId)
	args = append(args, user.Id)
	_, err = ar.db.Exec(ctx, query, args...)
	if err != nil {
		tx.Rollback(ctx)
		return newUsers, err
	}

	queryUsers := fmt.Sprintf("SELECT id, name, surname FROM %s", database.Client)
	rows, err := ar.db.Query(ctx, queryUsers)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}
	for rows.Next() {
		var user utils.UserDTO
		err := rows.Scan(user.Id, user.Name, user.Surname)
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

func (ar *AuthRepository) Delete(ctx context.Context, id int) (string, error) {
	query := fmt.Sprintf(`delete from %s where id=$1`, database.Client)
	_, err := ar.db.Exec(ctx, query, id)
	if err != nil {
		return "", err
	}
	return "User has been deleted", nil
}
