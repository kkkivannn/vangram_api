package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vangram_api/pkg/database"
	"vangram_api/utils"
)

type AuthRepositoryInterface interface {
	Create(user utils.UserDto) (int, error)
	Read(id int) (utils.UserDto, error)
	Update(id int, user utils.UserDto) (utils.UserDto, error)
	Delete(id int) (string, error)
}

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) Create(user utils.UserDto) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name) VALUES ($1) returning id", database.Client)
	row := ar.db.QueryRow(query, user.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (ar *AuthRepository) Read(id int) (utils.UserDto, error) {
	var user utils.UserDto
	query := fmt.Sprintf("select id, name from %s where id=$1", database.Client)
	if err := ar.db.Get(&user, query, id); err != nil {
		return utils.UserDto{}, err
	}
	return user, nil
}

func (ar *AuthRepository) Update(id int, user utils.UserDto) (utils.UserDto, error) {
	return utils.UserDto{}, nil
}

func (ar *AuthRepository) Delete(id int) (string, error) {
	var query = fmt.Sprintf("delete from %s where id=$1", database.Client)
	_, err := ar.db.Exec(query, id)
	if err != nil {
		return "", err
	}
	return "User was deleted", nil
}
