package service

import (
	//"crypto/sha1"
	//"fmt"
	//"github.com/dgrijalva/jwt-go"
	//todo "github.com/zmaxic1978/goweb"
	todo "github.com/zmaxic1978/goweb"
	"github.com/zmaxic1978/goweb/pkg/repository"
	//"time"
)

type ApiService struct {
	repo repository.Api
}

func NewApiService(repo repository.Api) *ApiService {
	return &ApiService{repo: repo}
}

// ----------------- Работа с авторами ----------------------

func (s *ApiService) CreateAuthor(author todo.Author) (int, error) {
	id, err := s.repo.CreateAuthor(author)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ApiService) GetAllAuthors() ([]todo.Author, error) {
	list, err := s.repo.GetAllAuthors()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *ApiService) GetAuthorById(id int) (todo.Author, error) {
	author, err := s.repo.GetAuthorById(id)
	if err != nil {
		return author, err
	}
	return author, nil
}

func (s *ApiService) SetAuthorById(author todo.Author) (int, error) {
	cnt, err := s.repo.SetAuthorById(author)
	if err != nil {
		return cnt, err
	}
	return cnt, nil
}

func (s *ApiService) DeleteAuthorById(authorId int) (int, error) {
	cnt, err := s.repo.DeleteAuthorById(authorId)
	if err != nil {
		return cnt, err
	}
	return cnt, nil
}

// ----------------- Работа с книгами -------------------------

func (s *ApiService) CreateBook(book todo.Book) (int, error) {
	id, err := s.repo.CreateBook(book)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ApiService) GetAllBooks() ([]todo.Book, error) {
	list, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *ApiService) GetBookById(id int) (todo.Book, error) {
	book, err := s.repo.GetBookById(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s *ApiService) SetBookById(book todo.Book) (int, error) {
	cnt, err := s.repo.SetBookById(book)
	if err != nil {
		return cnt, err
	}
	return cnt, nil
}

func (s *ApiService) DeleteBookById(bookId int) (int, error) {
	cnt, err := s.repo.DeleteBookById(bookId)
	if err != nil {
		return cnt, err
	}
	return cnt, nil
}

// ----------------- Работа с авторами и книгами -------------------------

func (s *ApiService) SetBookAuthorById(bookauthor todo.BookAuthor) (int, error) {
	cnt, err := s.repo.SetBookAuthorById(bookauthor)
	if err != nil {
		return cnt, err
	}
	return cnt, nil
}
