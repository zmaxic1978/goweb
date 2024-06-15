package repository

import (
	"database/sql"
	"fmt"
	"github.com/zmaxic1978/goweb/todo"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES($1, $2, $3) returning id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login todo.Login) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s where username = $1 and password_hash = $2", userTable)
	err := r.db.QueryRow(query, login.Username, login.Password).Scan(&user.Id)
	return user, err
}
