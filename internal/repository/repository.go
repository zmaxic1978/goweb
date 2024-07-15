package repository

import (
	"database/sql"
	todo2 "github.com/zmaxic1978/goweb/todo"
)

type Authorization interface {
	CreateUser(user todo2.User) (int, error)
	GetUser(login todo2.Login) (todo2.User, error)
}

type Api interface {
	// ----------------- Работа с авторами ----------------------
	CreateAuthor(author todo2.Author) (int, error)
	GetAllAuthors() ([]todo2.Author, error)
	GetAuthorById(id int) (todo2.Author, error)
	SetAuthorById(author todo2.Author) (int, error)
	DeleteAuthorById(authorId int) (int, error)
	// ----------------- Работа с книгами -------------------------
	CreateBook(book todo2.Book) (int, error)
	GetAllBooks() ([]todo2.Book, error)
	GetBookById(id int) (todo2.Book, error)
	SetBookById(book todo2.Book) (int, error)
	DeleteBookById(bookId int) (int, error)
	// ----------------- Работа с авторами и книгами -------------------------
	SetBookAuthorById(bookauthor todo2.BookAuthor) (int, error)
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
