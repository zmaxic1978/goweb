package service

import (
	todo "github.com/zmaxic1978/goweb"
	"github.com/zmaxic1978/goweb/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	CreateToken(login todo.Login) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	Api
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Api:           NewApiService(repos.Api),
	}
}
