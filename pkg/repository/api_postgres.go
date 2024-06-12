package repository

import (
	"database/sql"
	"errors"
	"fmt"
	todo "github.com/zmaxic1978/goweb"
	"strings"
)

const (
	errAuthorAlreadyExists = "похожий автор уже существует"
	errAuthorDoesntExists  = "автор с указанным id не существует"
	errBookAlreadyExists   = "похожая книга у автора уже существует"
	errBookDoesntExists    = "книга с указанным id не существует"
)

type ApiPostgres struct {
	db *sql.DB
	tm *TransactionManager
}

func NewApiPostgres(db *sql.DB) *ApiPostgres {
	return &ApiPostgres{
		db: db,
		tm: NewTransactionManager(db),
	}
}

// ----------------- Работа с авторами ----------------------

func (r *ApiPostgres) CreateAuthor(author todo.Author) (int, error) {
	// проверка на похожего автора
	authorId, err := r.sameAuthorExists(author)
	if err != nil {
		return authorId, err
	}

	// добавляем нового автора
	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname, description, birthday) VALUES($1, $2, $3, $4) returning id", authorsTable)
	row := r.tm.QueryRow(query, author.FirstName, author.LastName, author.Description, author.Birthday)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ApiPostgres) GetAllAuthors() ([]todo.Author, error) {
	var list []todo.Author
	// выбираем всех авторов
	query := fmt.Sprintf("SELECT id, firstname, lastname, description, birthday FROM %s", authorsTable)
	rows, err := r.tm.Query(query)
	if err != nil {
		return nil, todo.DBError{Message: err.Error()}
	}

	defer rows.Close()
	for rows.Next() {
		var a todo.Author
		if err := rows.Scan(&a.Id, &a.FirstName, &a.LastName, &a.Description, &a.Birthday); err != nil {
			return nil, err
		}
		list = append(list, a)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *ApiPostgres) GetAuthorById(id int) (todo.Author, error) {
	var a todo.Author

	// проверка на существующего автора
	_, err := r.authorExists(id)
	if err != nil {
		return a, err
	}

	// выбираем нужного автора
	query := fmt.Sprintf("SELECT id, firstname, lastname, description, birthday FROM %s WHERE id = $1", authorsTable)
	row := r.tm.QueryRow(query, id)
	err = row.Scan(&a.Id, &a.FirstName, &a.LastName, &a.Description, &a.Birthday)
	if err == sql.ErrNoRows {
		return a, todo.NoDataFound{Message: errAuthorDoesntExists}
	}
	if err != nil {
		return a, todo.DBError{Message: err.Error()}
	}
	return a, nil
}

func (r *ApiPostgres) SetAuthorById(author todo.Author) (int, error) {
	// проверка на существующего автора
	_, err := r.authorExists(author.Id)
	if err != nil {
		return 0, err
	}

	// проверка на похожего автора
	authorId, err := r.sameAuthorExists(author)
	if err != nil {
		return authorId, err
	}

	// обновляем информацию по автору
	query := fmt.Sprintf(""+
		"UPDATE %s "+
		"SET firstname = $1, lastname = $2, description = $3, birthday = $4 "+
		"WHERE id = $5", authorsTable)
	res, err := r.tm.Exec(query, author.FirstName, author.LastName, author.Description, author.Birthday, author.Id)
	if err != nil {
		return 0, todo.DBError{Message: err.Error()}
	}
	ra, err2 := res.RowsAffected()
	if err2 != nil {
		return 0, err2
	}

	return int(ra), nil
}

func (r *ApiPostgres) DeleteAuthorById(authorId int) (int, error) {
	// проверка на существующего автора
	_, err := r.authorExists(authorId)
	if err != nil {
		return 0, err
	}

	// удаляем автора
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", authorsTable)
	res, err := r.tm.Exec(query, authorId)
	if err != nil {
		return 0, todo.DBError{Message: err.Error()}
	}
	ra, err2 := res.RowsAffected()
	if err2 != nil {
		return 0, err2
	}

	return int(ra), nil
}

// ----------------- Работа с книгами -------------------------

func (r *ApiPostgres) CreateBook(book todo.Book) (int, error) {

	// проверка на существующего автора
	_, err := r.authorExists(book.AuthorId)
	if err != nil {
		return 0, err
	}

	// проверка на похожую книгу
	bookId, err := r.sameBookExists(book)
	if err != nil {
		return bookId, err
	}

	// добавляем книгу
	query := fmt.Sprintf("INSERT INTO %s (name, authorid, year, isbn) VALUES($1, $2, $3, $4) returning id", booksTable)
	row := r.tm.QueryRow(query, book.Name, book.AuthorId, book.Year, book.ISBN)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ApiPostgres) GetAllBooks() ([]todo.Book, error) {
	var list []todo.Book
	// выбираем все книги
	query := fmt.Sprintf("SELECT id, name, authorid, year, isbn FROM %s", booksTable)
	rows, err := r.tm.Query(query)
	if err != nil {
		return nil, todo.DBError{Message: err.Error()}
	}

	defer rows.Close()
	for rows.Next() {
		var a todo.Book
		if err := rows.Scan(&a.Id, &a.Name, &a.AuthorId, &a.Year, &a.ISBN); err != nil {
			return nil, err
		}
		list = append(list, a)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *ApiPostgres) GetBookById(id int) (todo.Book, error) {
	var book todo.Book

	// проверка на существующую книгу
	_, err := r.bookExists(id)
	if err != nil {
		return book, err
	}

	// выбираем нужную книгу
	query := fmt.Sprintf("SELECT id, name, authorid, year, isbn FROM %s WHERE id = $1", booksTable)
	row := r.tm.QueryRow(query, id)
	err = row.Scan(&book.Id, &book.Name, &book.AuthorId, &book.Year, &book.ISBN)
	if err == sql.ErrNoRows {
		return book, todo.NoDataFound{Message: errBookDoesntExists}
	}
	if err != nil {
		return book, todo.DBError{Message: err.Error()}
	}
	return book, nil
}

func (r *ApiPostgres) SetBookById(book todo.Book) (int, error) {
	// проверка на существующую книгу
	_, err := r.bookExists(book.Id)
	if err != nil {
		return 0, err
	}

	// проверка на существующего автора
	_, err = r.authorExists(book.AuthorId)
	if err != nil {
		return 0, err
	}

	// проверка на похожую книгу
	bookId, err := r.sameBookExists(book)
	if err != nil {
		return bookId, err
	}

	// обновляем информацию по книге
	query := fmt.Sprintf(""+
		"UPDATE %s "+
		"SET name = $1, authorid = $2, year = $3, isbn = $4 "+
		"WHERE id = $5", booksTable)
	res, err := r.tm.Exec(query, book.Name, book.AuthorId, book.Year, book.ISBN, book.Id)
	if err != nil {
		return 0, todo.DBError{Message: err.Error()}
	}

	ra, err2 := res.RowsAffected()
	if err2 != nil {
		return 0, err2
	}

	return int(ra), nil
}

func (r *ApiPostgres) DeleteBookById(bookId int) (int, error) {
	// проверка на существующую книгу
	_, err := r.bookExists(bookId)
	if err != nil {
		return 0, err
	}

	// удаляем книгу
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", booksTable)
	res, err := r.tm.Exec(query, bookId)
	if err != nil {
		return 0, todo.DBError{Message: err.Error()}
	}
	ra, err2 := res.RowsAffected()
	if err2 != nil {
		return 0, err2
	}

	return int(ra), nil
}

// ----------------- Работа с авторами и книгами -------------------------

func (r *ApiPostgres) SetBookAuthorById(bookauthor todo.BookAuthor) (int, error) {

	// начинаем транзакцию
	err := r.tm.StartTransaction()
	if err != nil {
		return 0, err
	}

	_, err = r.SetBookById(bookauthor.Book)
	err = errors.New("bla bla bla")
	if err != nil {
		r.tm.RollBack()
		return 0, err
	}

	_, err = r.SetAuthorById(bookauthor.Author)
	if err != nil {
		r.tm.RollBack()
		return 0, err
	}

	err = r.tm.Commit()
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// ----------------- Вспомогательные функции -----------------------------

func (r *ApiPostgres) authorExists(authorId int) (int, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE id = $1 ", authorsTable)
	row := r.tm.QueryRow(query, authorId)
	var id int
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return 0, todo.NoDataFound{Message: errAuthorDoesntExists}
	}
	if err != nil && !(errors.Is(err, sql.ErrNoRows)) {
		return 0, todo.DBError{Message: err.Error()}
	}

	return id, nil
}

func (r *ApiPostgres) bookExists(bookId int) (int, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE id = $1 ", booksTable)
	row := r.tm.QueryRow(query, bookId)
	var id int
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return 0, todo.NoDataFound{Message: errBookDoesntExists}
	}
	if err != nil && !(errors.Is(err, sql.ErrNoRows)) {
		return 0, todo.DBError{Message: err.Error()}
	}

	return id, nil
}

func (r *ApiPostgres) sameAuthorExists(author todo.Author) (int, error) {
	// проверка на похожую книгу у автора
	query := fmt.Sprintf(""+
		"SELECT id "+
		" FROM %s "+
		"WHERE UPPER(firstname) LIKE CONCAT('%%',$1::text,'%%') "+
		"  AND UPPER(lastname) LIKE CONCAT('%%',$2::text,'%%') "+
		"  AND birthday = $3 "+
		"  AND id <> $4", authorsTable)
	row := r.tm.QueryRow(query, strings.ToUpper(author.FirstName), strings.ToUpper(author.LastName), strings.ToUpper(author.Birthday), author.Id)
	var id int
	err := row.Scan(&id)
	if err == nil && id > 0 {
		return id, todo.BadFormatError{Message: errAuthorAlreadyExists}
	}
	if err != nil && !(errors.Is(err, sql.ErrNoRows)) {
		return 0, err
	}

	return id, nil
}

func (r *ApiPostgres) sameBookExists(book todo.Book) (int, error) {
	// проверка на похожую книгу у автора
	query := fmt.Sprintf(""+
		"SELECT id "+
		" FROM %s "+
		"WHERE UPPER(name) LIKE CONCAT('%%',$1::text,'%%') "+
		"  AND year = $2 "+
		"  AND UPPER(isbn) LIKE CONCAT('%%',$3::text,'%%') "+
		"  AND id <> $4", booksTable)
	row := r.tm.QueryRow(query, strings.ToUpper(book.Name), book.Year, strings.ToUpper(book.ISBN), book.Id)
	var id int
	err := row.Scan(&id)
	if err == nil && id > 0 {
		return id, todo.BadFormatError{Message: errBookAlreadyExists}
	}
	if err != nil && !(errors.Is(err, sql.ErrNoRows)) {
		return 0, err
	}

	return id, nil
}
