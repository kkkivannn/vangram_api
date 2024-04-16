package storage

import (
	"context"
	"errors"
	"fmt"
	"vangram_api/internal/service/user"
	"vangram_api/internal/storage/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(ctx context.Context, user user.SaveUser) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, age, phone_number, photo, created_at, uploaded_at) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id", postgres.User)
	row := s.db.QueryRow(ctx, query, user.Name, user.Surname, user.Age, user.Phone, user.Photo, user.CreatedAt, user.UploadedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserStorage) ReadUser(ctx context.Context, id int) (user.User, error) {
	var resp user.User
	query := fmt.Sprintf("SELECT id, name, surname, age, phone_number, photo FROM %s WHERE id=$1", postgres.User)
	err := s.db.QueryRow(ctx, query, id).Scan(&resp.ID, &resp.Name, &resp.Surname, &resp.Age, &resp.Phone, &resp.Photo)
	if err != nil {
		return user.User{}, err
	}
	if resp.Photo != nil {
		photo := fmt.Sprintf("%s%s", url, *resp.Photo)
		resp.Photo = &photo
	}

	return resp, nil
}

func (s *UserStorage) UpdateUser(ctx context.Context, reqUser user.SaveUser, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, surname=$2, age=$3, photo=$4, uploaded_at=$5 WHERE id=$6", postgres.User)
	data, err := s.db.Exec(ctx, query, reqUser.Name, reqUser.Surname, reqUser.Age, reqUser.Photo, reqUser.UploadedAt, userId)
	if err != nil {
		return err
	}
	if num := data.RowsAffected(); num == 0 {
		return errors.New("updated 0 rows")
	}
	return nil

	//tx, err := s.db.Begin(ctx)
	//if err != nil {ф
	//	return user.UpdatedUser{}, err
	//}
	//setValues := make([]string, 0)
	//args := make([]interface{}, 0)
	//argId := 1
	//var newUsers user.User
	//if reqUser.Name != nil {
	//	setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
	//	args = append(args, *reqUser.Name)
	//	argId++
	//}
	//if reqUser.Surname != nil {
	//	setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
	//	args = append(args, *u.Surname)
	//	argId++
	//}

}

func (s *UserStorage) DeleteUser(ctx context.Context, id int) (string, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, postgres.User)
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return "", err
	}
	return "User has been deleted", nil
}

func (s *UserStorage) GetAllUsers(ctx context.Context) ([]user.User, error) {
	var users []user.User
	query := fmt.Sprintf("SELECT id, name, surname FROM %s", postgres.User)
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

func (s *UserStorage) ReadUserByNumber(ctx context.Context, number string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE phone_number=$1", postgres.User)
	err := s.db.QueryRow(ctx, query, number).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserStorage) SetUserRefreshToken(ctx context.Context, token string, userId int, sessionId string) error {
	query := fmt.Sprintf("INSERT INTO %s (refresh_token, user_id, session_id) VALUES ($1, $2, $3)", postgres.UsersTokens)
	_, err := s.db.Exec(ctx, query, token, userId, sessionId)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStorage) UpdateUserSession(ctx context.Context, sessionID string, userID int, oldRefreshToken string, newRefreshToken string) error {
	query := fmt.Sprintf("UPDATE %s SET refresh_token=$1 WHERE refresh_token=$2 AND session_id=$3 AND user_id=$4", postgres.UsersTokens)
	data, err := s.db.Exec(ctx, query, newRefreshToken, oldRefreshToken, sessionID, userID)
	if err != nil {
		return err
	}
	if num := data.RowsAffected(); num == 0 {
		return errors.New("updated 0 rows")
	}
	return nil
}

func (s *UserStorage) DeleteUserSession(ctx context.Context, sessionId string, userID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE session_id=$1 AND user_id=$2", postgres.UsersTokens)
	data, err := s.db.Exec(ctx, query, sessionId, userID)
	if err != nil {
		return err
	}
	if num := data.RowsAffected(); num == 0 {
		return errors.New("deleted 0 rows")
	}
	return nil
}
