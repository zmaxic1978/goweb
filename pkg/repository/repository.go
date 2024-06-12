package repository

import (
	"database/sql"
	todo "github.com/zmaxic1978/goweb"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(login todo.Login) (todo.User, error)
}

type Api interface {
	// ----------------- Работа с авторами ----------------------
	CreateAuthor(author todo.Author) (int, error)
	GetAllAuthors() ([]todo.Author, error)
	GetAuthorById(id int) (todo.Author, error)
	SetAuthorById(author todo.Author) (int, error)
	DeleteAuthorById(authorId int) (int, error)
	// ----------------- Работа с книгами -------------------------
	CreateBook(book todo.Book) (int, error)
	GetAllBooks() ([]todo.Book, error)
	GetBookById(id int) (todo.Book, error)
	SetBookById(book todo.Book) (int, error)
	DeleteBookById(bookId int) (int, error)
	// ----------------- Работа с авторами и книгами -------------------------
	SetBookAuthorById(bookauthor todo.BookAuthor) (int, error)
}

type Transaction interface {
	StartTransaction() error
	Commit() error
	RollBack() error
	QueryRow(query string, args ...any) *sql.Row
	Query(query string, args ...any) (*sql.Rows, error)
	ExecContext(query string, args ...any) (sql.Result, error)
	PrepareContext(query string) (*sql.Stmt, error)
}

type Repository struct {
	Authorization
	Api
	Transaction
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Api:           NewApiPostgres(db),
	}
}
