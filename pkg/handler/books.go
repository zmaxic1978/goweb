package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	todo2 "github.com/zmaxic1978/goweb/todo"
	"net/http"
	"strconv"
)

const (
	errIncorrectBookFormat = "Неверное передана информация по книге"
	errInvalidBookId       = "передан неверный Id книги"
)

type responseGetAllBooks struct {
	Data []todo2.Book `json:"data"`
}

func (h *Handler) createBook(c *gin.Context) {

	/*if _, err := getUserId(c); err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}*/

	var input todo2.Book
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errIncorrectBookFormat, err).Error())
		return
	}

	id, err := h.services.Api.CreateBook(input)
	if err != nil {
		if errors.As(err, new(todo2.BadFormatError)) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) getAllBooks(c *gin.Context) {
	list, err := h.services.Api.GetAllBooks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responseGetAllBooks{Data: list})
}

func (h *Handler) getBookById(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidBookId, err).Error())
		return
	}

	book, err := h.services.Api.GetBookById(bookId)
	if err != nil && errors.As(err, new(todo2.NoDataFound)) {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) setBookById(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidBookId, err).Error())
		return
	}

	var input todo2.Book
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errIncorrectBookFormat, err).Error())
		return
	}

	input.Id = bookId
	cnt, err := h.services.Api.SetBookById(input)
	if err != nil {
		if errors.As(err, new(todo2.BadFormatError)) || errors.As(err, new(todo2.NoDataFound)) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"записей обновлено": cnt})
}

func (h *Handler) deleteBookById(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidBookId, err).Error())
		return
	}

	cnt, err := h.services.Api.DeleteBookById(bookId)
	if err != nil {
		if errors.As(err, new(todo2.BadFormatError)) || errors.As(err, new(todo2.NoDataFound)) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"записей удалено": cnt})
}
