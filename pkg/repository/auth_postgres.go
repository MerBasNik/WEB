package repository

import (
	"fmt"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
	//"github.com/pelletier/go-toml/query"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user chat.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash, token) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Email, user.Password, "")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (chat.User, error) {
	var user chat.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *AuthPostgres) GetUserEmail(token string) (chat.User, error) {
	var user chat.User
	query := fmt.Sprintf("SELECT tl.email FROM %s tl WHERE tl.token=$1", usersTable)
	err := r.db.Get(&user, query, token)

	return user, err
}

func (r *AuthPostgres) ResetPassword(email, password string) error {
	query := fmt.Sprintf("UPDATE %s tl SET password_hash=$1 WHERE tl.email=$2", usersTable)
	
	_, err := r.db.Exec(query, password, email)

	return err
}

func (r *AuthPostgres) DeleteUserToken(user chat.User) error {
	query := fmt.Sprintf("UPDATE %s tl SET token=$1 WHERE tl.email=$2", usersTable)
	
	_, err := r.db.Exec(query, "", user.Email)

	return err
}

func (r *AuthPostgres) SetUserToken(token, email string) error {
	query := fmt.Sprintf("UPDATE %s tl SET token=$1 WHERE tl.email=$2", usersTable)
	
	_, err := r.db.Exec(query, token, email)

	return err
}