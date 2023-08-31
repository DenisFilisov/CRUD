package repository

import (
	"CRUD/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (r *AuthPostgres) GetUserById(userId int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, userId)
	return user, err
}

func (r *AuthPostgres) CheckRefreshToken(refreshToken string) (int, error) {
	var rt model.RefreshToken
	query := fmt.Sprintf("SELECT * FROM %s WHERE token=$1", refreshTokenTable)
	err := r.db.Get(&rt, query, refreshToken)
	return rt.UserID, err
}

func (r *AuthPostgres) SaveRefreshToken(oldToken, refreshToken string, userId int) error {
	refreshTokenTTL := time.Duration(viper.GetInt("tokens.refreshTokenTTl"))

	var token model.RefreshToken
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 and token=$2", refreshTokenTable)
	err := r.db.Get(&token, query, userId, oldToken)
	if err == nil {
		insert := fmt.Sprintf("UPDATE %s SET token=$1, expires_at=$2 WHERE user_id=$3 and token=$4", refreshTokenTable)
		row := r.db.QueryRow(insert, refreshToken, time.Now().Add(refreshTokenTTL*time.Hour), userId, oldToken)
		return row.Err()
	} else {
		insert := fmt.Sprintf("INSERT INTO %s (user_id, token, expires_at) values ($1, $2, $3)", refreshTokenTable)
		row := r.db.QueryRow(insert, userId, refreshToken, time.Now().Add(refreshTokenTTL*time.Hour))
		return row.Err()
	}
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int

	insert := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(insert, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) FindUserByUserNameAndPswd(username, password string) (model.User, error) {
	var user model.User
	insert := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, insert, username, password)
	return user, err
}
