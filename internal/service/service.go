package service

import (
	"github.com/zmaxic1978/goweb/internal/repository"
	todo2 "github.com/zmaxic1978/goweb/todo"
)

type Authorization interface {
	CreateUser(user todo2.User) (int, error)
	CreateToken(login todo2.Login) (string, error)
	ParseToken(token string) (int, error)
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
