package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/zmaxic1978/goweb"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(login todo.Login) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
